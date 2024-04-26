package repo

import (
	"context"
	"testing"

	"github.com/jae2274/careerhub-userinfo-service/careerhub/userinfo_service/common/domain/matchjob"
	rRepo "github.com/jae2274/careerhub-userinfo-service/careerhub/userinfo_service/restapi/repo"
	sRepo "github.com/jae2274/careerhub-userinfo-service/careerhub/userinfo_service/suggester/repo"
	"github.com/jae2274/careerhub-userinfo-service/test/tinit"
	"github.com/jae2274/goutils/ptr"
	"github.com/stretchr/testify/require"
)

func TestMatchJobRepo(t *testing.T) {

	t.Run("return empty conditions if nothing saved", func(t *testing.T) {
		ctx := context.Background()
		suggesterRepo := sRepo.NewMatchJobRepo(tinit.InitDB(t))

		desireCondition, err := suggesterRepo.GetMatchJobs(ctx)
		require.NoError(t, err)

		require.Empty(t, desireCondition)
	})

	t.Run("return empty if all agreeToMail false", func(t *testing.T) {
		ctx := context.Background()
		suggesterRepo := sRepo.NewMatchJobRepo(tinit.InitDB(t))
		restapiRepo := rRepo.NewMatchJobRepo(tinit.InitDB(t))

		isSuccess, err := restapiRepo.InitMatchJob(ctx, "test_user_id")
		require.NoError(t, err)
		require.True(t, isSuccess)

		isSuccess, err = restapiRepo.InitMatchJob(ctx, "test_user_id2")
		require.NoError(t, err)
		require.True(t, isSuccess)

		desireCondition, err := suggesterRepo.GetMatchJobs(ctx)
		require.NoError(t, err)

		require.Empty(t, desireCondition)
	})

	t.Run("return matchjobs if all agreeToMail true", func(t *testing.T) {
		ctx := context.Background()
		suggesterRepo := sRepo.NewMatchJobRepo(tinit.InitDB(t))
		restapiRepo := rRepo.NewMatchJobRepo(tinit.InitDB(t))

		restapiRepo.InitMatchJob(ctx, "test_user_id")
		restapiRepo.InitMatchJob(ctx, "test_user_id2")

		restapiRepo.UpdateAgreeToMail(ctx, "test_user_id", true)
		restapiRepo.UpdateAgreeToMail(ctx, "test_user_id2", true)

		desireCondition, err := suggesterRepo.GetMatchJobs(ctx)
		require.NoError(t, err)

		require.Len(t, desireCondition, 2)
		require.Equal(t, "test_user_id", desireCondition[0].UserId)
		require.Equal(t, "test_user_id2", desireCondition[1].UserId)
	})

	t.Run("return same matchjobs", func(t *testing.T) {
		ctx := context.Background()
		suggesterRepo := sRepo.NewMatchJobRepo(tinit.InitDB(t))
		restapiRepo := rRepo.NewMatchJobRepo(tinit.InitDB(t))

		isSuccess, err := restapiRepo.InitMatchJob(ctx, "test_user_id")
		require.NoError(t, err)
		require.True(t, isSuccess)

		isSuccess, err = restapiRepo.UpdateAgreeToMail(ctx, "test_user_id", true)
		require.NoError(t, err)
		require.True(t, isSuccess)

		savedConditions := []*matchjob.Condition{newCondition(), newCondition()}
		updatedConditions := make([]*matchjob.Condition, len(savedConditions))

		for i, c := range savedConditions {
			isSuccess, err := restapiRepo.InsertCondition(ctx, "test_user_id", 2, c)
			require.NoError(t, err)
			require.True(t, isSuccess)

			updatedConditions[i] = newUpdatedCondition(c.ConditionId)
			isSuccess, err = restapiRepo.UpdateCondition(ctx, "test_user_id", updatedConditions[i])
			require.NoError(t, err)
			require.True(t, isSuccess)
		}

		desireCondition, err := suggesterRepo.GetMatchJobs(ctx)
		require.NoError(t, err)

		require.Len(t, desireCondition, 1)
		require.Equal(t, "test_user_id", desireCondition[0].UserId)
		require.Len(t, desireCondition[0].Conditions, 2)
		require.Equal(t, updatedConditions, desireCondition[0].Conditions)
	})
}

func newCondition() *matchjob.Condition {
	return &matchjob.Condition{
		ConditionName: "conditionName",
		Query: &matchjob.Query{
			Categories: []*matchjob.CategoryQuery{
				{Site: "jumpit", CategoryName: "backend"},
				{Site: "wanted", CategoryName: "backend"},
			},
			SkillNames: [][]string{{"skillName"}},
			MinCareer:  ptr.P(int32(1)),
			MaxCareer:  ptr.P(int32(2)),
		},
	}
}

func newUpdatedCondition(conditionId string) *matchjob.Condition {
	return &matchjob.Condition{
		ConditionId:   conditionId,
		ConditionName: "updatedConditionName",
		Query: &matchjob.Query{
			Categories: []*matchjob.CategoryQuery{
				{Site: "jumpit", CategoryName: "frontend"},
				{Site: "wanted", CategoryName: "frontend"},
			},
			SkillNames: [][]string{{"updatedSkillName"}},
			MinCareer:  ptr.P(int32(3)),
			MaxCareer:  ptr.P(int32(5)),
		},
	}
}
