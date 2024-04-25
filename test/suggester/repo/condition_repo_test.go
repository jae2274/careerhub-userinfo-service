package repo

import (
	"context"
	"testing"

	"github.com/jae2274/careerhub-userinfo-service/careerhub/userinfo_service/common/domain/condition"
	rRepo "github.com/jae2274/careerhub-userinfo-service/careerhub/userinfo_service/restapi/repo"
	sRepo "github.com/jae2274/careerhub-userinfo-service/careerhub/userinfo_service/suggester/repo"
	"github.com/jae2274/careerhub-userinfo-service/test/tinit"
	"github.com/jae2274/goutils/ptr"
	"github.com/stretchr/testify/require"
)

func TestConditionRepo(t *testing.T) {

	t.Run("return empty conditions if nothing saved", func(t *testing.T) {
		ctx := context.Background()
		suggesterRepo := sRepo.NewConditionRepo(tinit.InitDB(t))

		desireCondition, err := suggesterRepo.GetDesiredConditions(ctx)
		require.NoError(t, err)

		require.Empty(t, desireCondition)
	})

	t.Run("return empty if all agreeToMail false", func(t *testing.T) {
		ctx := context.Background()
		suggesterRepo := sRepo.NewConditionRepo(tinit.InitDB(t))
		restapiRepo := rRepo.NewConditionRepo(tinit.InitDB(t))

		isSuccess, err := restapiRepo.InitDesiredCondition(ctx, "test_user_id")
		require.NoError(t, err)
		require.True(t, isSuccess)

		isSuccess, err = restapiRepo.InitDesiredCondition(ctx, "test_user_id2")
		require.NoError(t, err)
		require.True(t, isSuccess)

		desireCondition, err := suggesterRepo.GetDesiredConditions(ctx)
		require.NoError(t, err)

		require.Empty(t, desireCondition)
	})

	t.Run("return desired conditions if all agreeToMail true", func(t *testing.T) {
		ctx := context.Background()
		suggesterRepo := sRepo.NewConditionRepo(tinit.InitDB(t))
		restapiRepo := rRepo.NewConditionRepo(tinit.InitDB(t))

		restapiRepo.InitDesiredCondition(ctx, "test_user_id")
		restapiRepo.InitDesiredCondition(ctx, "test_user_id2")

		restapiRepo.UpdateAgreeToMail(ctx, "test_user_id", true)
		restapiRepo.UpdateAgreeToMail(ctx, "test_user_id2", true)

		desireCondition, err := suggesterRepo.GetDesiredConditions(ctx)
		require.NoError(t, err)

		require.Len(t, desireCondition, 2)
		require.Equal(t, "test_user_id", desireCondition[0].UserId)
		require.Equal(t, "test_user_id2", desireCondition[1].UserId)
	})

	t.Run("return same desired conditions", func(t *testing.T) {
		ctx := context.Background()
		suggesterRepo := sRepo.NewConditionRepo(tinit.InitDB(t))
		restapiRepo := rRepo.NewConditionRepo(tinit.InitDB(t))

		isSuccess, err := restapiRepo.InitDesiredCondition(ctx, "test_user_id")
		require.NoError(t, err)
		require.True(t, isSuccess)

		isSuccess, err = restapiRepo.UpdateAgreeToMail(ctx, "test_user_id", true)
		require.NoError(t, err)
		require.True(t, isSuccess)

		savedConditions := []*condition.Condition{newCondition(), newCondition()}
		updatedConditions := make([]*condition.Condition, len(savedConditions))

		for i, c := range savedConditions {
			isSuccess, err := restapiRepo.InsertCondition(ctx, "test_user_id", 2, c)
			require.NoError(t, err)
			require.True(t, isSuccess)

			updatedConditions[i] = newUpdatedCondition(c.ConditionId)
			isSuccess, err = restapiRepo.UpdateCondition(ctx, "test_user_id", updatedConditions[i])
			require.NoError(t, err)
			require.True(t, isSuccess)
		}

		desireCondition, err := suggesterRepo.GetDesiredConditions(ctx)
		require.NoError(t, err)

		require.Len(t, desireCondition, 1)
		require.Equal(t, "test_user_id", desireCondition[0].UserId)
		require.Len(t, desireCondition[0].Conditions, 2)
		require.Equal(t, updatedConditions, desireCondition[0].Conditions)
	})
}

func newCondition() *condition.Condition {
	return &condition.Condition{
		ConditionName: "conditionName",
		Query: &condition.Query{
			Categories: []*condition.CategoryQuery{
				{Site: "jumpit", CategoryName: "backend"},
				{Site: "wanted", CategoryName: "backend"},
			},
			SkillNames: [][]string{{"skillName"}},
			MinCareer:  ptr.P(int32(1)),
			MaxCareer:  ptr.P(int32(2)),
		},
	}
}

func newUpdatedCondition(conditionId string) *condition.Condition {
	return &condition.Condition{
		ConditionId:   conditionId,
		ConditionName: "updatedConditionName",
		Query: &condition.Query{
			Categories: []*condition.CategoryQuery{
				{Site: "jumpit", CategoryName: "frontend"},
				{Site: "wanted", CategoryName: "frontend"},
			},
			SkillNames: [][]string{{"updatedSkillName"}},
			MinCareer:  ptr.P(int32(3)),
			MaxCareer:  ptr.P(int32(5)),
		},
	}
}
