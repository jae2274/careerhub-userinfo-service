package suggester

import (
	"context"
	"fmt"
	"testing"

	"github.com/jae2274/careerhub-userinfo-service/careerhub/userinfo_service/restapi/restapi_grpc"
	"github.com/jae2274/careerhub-userinfo-service/test/tinit"
	"github.com/jae2274/goutils/ptr"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestSuggesterGrpc(t *testing.T) {
	cancelFunc := tinit.RunTestApp(t)
	defer cancelFunc()

	t.Run("returm empty", func(t *testing.T) {
		ctx := context.Background()
		client := tinit.InitSuggesterClient(t)
		conditions, err := client.GetConditions(ctx, &emptypb.Empty{})
		require.NoError(t, err)
		require.Len(t, conditions.Conditions, 0)
	})

	t.Run("return conditions", func(t *testing.T) {
		tinit.InitDB(t)
		ctx := context.Background()

		restapiClient := tinit.InitMatchJobGrpcClient(t)
		updateAgreeToMail := func(ctx context.Context, req *restapi_grpc.UpdateAgreeToMailRequest) {
			isSuccess, err := restapiClient.UpdateAgreeToMail(ctx, req)
			require.NoError(t, err)
			require.True(t, isSuccess.IsSuccess)
		}
		addCondition := func(ctx context.Context, req *restapi_grpc.AddConditionRequest) {
			isSuccess, err := restapiClient.AddCondition(ctx, req)
			require.NoError(t, err)
			require.True(t, isSuccess.IsSuccess)
		}
		userId1 := "user1"
		userId2 := "user2"

		requests := []*restapi_grpc.AddConditionRequest{
			newAddConditionRequest(userId1, 1),
			newAddConditionRequest(userId1, 2),
			newAddConditionRequest(userId2, 3),
			newAddConditionRequest(userId2, 4),
		}

		for _, req := range requests {
			addCondition(ctx, req)
		}

		updateAgreeToMail(ctx, &restapi_grpc.UpdateAgreeToMailRequest{UserId: userId1, AgreeToMail: true})
		updateAgreeToMail(ctx, &restapi_grpc.UpdateAgreeToMailRequest{UserId: userId2, AgreeToMail: false})

		suggesterClient := tinit.InitSuggesterClient(t)
		res, err := suggesterClient.GetConditions(ctx, &emptypb.Empty{})
		require.NoError(t, err)
		require.Len(t, res.Conditions, 2)

		for i, req := range requests[0:2] {
			resCondition := res.Conditions[i]
			require.Equal(t, userId1, resCondition.UserId)
			require.Equal(t, req.Condition.ConditionName, resCondition.ConditionName)
			require.Equal(t, req.Condition.Query.MinCareer, resCondition.Query.MinCareer)
			require.Equal(t, req.Condition.Query.MaxCareer, resCondition.Query.MaxCareer)

			require.Len(t, resCondition.Query.Categories, len(req.Condition.Query.Categories))
			for j, category := range req.Condition.Query.Categories {
				require.Equal(t, category.Site, resCondition.Query.Categories[j].Site)
				require.Equal(t, category.CategoryName, resCondition.Query.Categories[j].CategoryName)
			}

			require.Len(t, resCondition.Query.SkillNames, len(req.Condition.Query.SkillNames))
			for j, skill := range req.Condition.Query.SkillNames {
				require.Equal(t, skill.Or, resCondition.Query.SkillNames[j].Or)
			}
		}
	})
}

func newAddConditionRequest(userId string, number int) *restapi_grpc.AddConditionRequest {
	attachN := func(str string, n int) string {
		return fmt.Sprintf("%s%d", str, n)
	}
	return &restapi_grpc.AddConditionRequest{
		UserId:     userId,
		LimitCount: uint32(2),
		Condition: &restapi_grpc.AddConditionReq{
			ConditionName: attachN("ConditionName", number),
			Query: &restapi_grpc.Query{
				Categories: []*restapi_grpc.Category{{Site: "Site", CategoryName: "CategoryName"}},
				SkillNames: []*restapi_grpc.Skill{{Or: []string{"SkillName"}}},
				MinCareer:  ptr.P(int32(1)),
				MaxCareer:  ptr.P(int32(3)),
			},
		},
	}
}
