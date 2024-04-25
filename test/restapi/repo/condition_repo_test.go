package repo

import (
	"context"
	"testing"
	"time"

	"github.com/jae2274/careerhub-userinfo-service/careerhub/userinfo_service/common/domain/condition"
	"github.com/jae2274/careerhub-userinfo-service/careerhub/userinfo_service/restapi/repo"
	"github.com/jae2274/careerhub-userinfo-service/test/tinit"
	"github.com/jae2274/goutils/ptr"
	"github.com/stretchr/testify/require"
)

func TestConditionRepo(t *testing.T) {
	userId := "test_user_id"

	t.Run("FindByUserId without InitConditions", func(t *testing.T) {
		conditionRepo := initConditionRepo(t)
		ctx := context.TODO()

		desiredCondition, err := conditionRepo.FindByUserId(ctx, userId)
		require.NoError(t, err)
		require.Nil(t, desiredCondition)
	})

	t.Run("FindByUserId after InitConditions", func(t *testing.T) {
		conditionRepo := initConditionRepo(t)
		ctx := context.TODO()

		success, err := conditionRepo.InitDesiredCondition(ctx, userId)
		require.NoError(t, err)
		require.True(t, success)

		desiredCondition, err := conditionRepo.FindByUserId(ctx, userId)
		require.NoError(t, err)
		require.NotNil(t, desiredCondition)
		require.Equal(t, userId, desiredCondition.UserId)
		require.Len(t, desiredCondition.Conditions, 0)

		now := time.Now()
		checkSimilarTimes(t, now, desiredCondition.InsertedAt)
	})

	t.Run("FindByUserId without InsertCondition", func(t *testing.T) {
		conditionRepo := initConditionRepo(t)
		ctx := context.TODO()

		success, err := conditionRepo.InitDesiredCondition(ctx, userId)
		require.NoError(t, err)
		require.True(t, success)

		finded, err := conditionRepo.FindByUserId(ctx, userId)
		require.NoError(t, err)
		require.Len(t, finded.Conditions, 0)
	})

	t.Run("FindByUserId after InsertCondition", func(t *testing.T) {
		conditionRepo := initConditionRepo(t)
		ctx := context.TODO()

		success, err := conditionRepo.InitDesiredCondition(ctx, userId)
		require.NoError(t, err)
		require.True(t, success)

		savedCondition := newCondition()
		success, err = conditionRepo.InsertCondition(ctx, userId, 1, savedCondition)
		require.NoError(t, err)
		require.True(t, success)

		finded, err := conditionRepo.FindByUserId(ctx, userId)
		require.NoError(t, err)
		require.Len(t, finded.Conditions, 1)
		require.Equal(t, savedCondition, finded.Conditions[0])
	})

	t.Run("FindByUserId after UpdateCondition", func(t *testing.T) {
		conditionRepo := initConditionRepo(t)
		ctx := context.TODO()

		success, err := conditionRepo.InitDesiredCondition(ctx, userId)
		require.NoError(t, err)
		require.True(t, success)

		savedCondition := newCondition()
		success, err = conditionRepo.InsertCondition(ctx, userId, 1, savedCondition)
		require.NoError(t, err)
		require.True(t, success)

		updatedCondition := newUpdatedCondition(savedCondition.ConditionId)
		success, err = conditionRepo.UpdateCondition(ctx, userId, updatedCondition)
		require.NoError(t, err)
		require.True(t, success)

		finded, err := conditionRepo.FindByUserId(ctx, userId)
		require.NoError(t, err)
		require.Len(t, finded.Conditions, 1)
		require.NotEqual(t, savedCondition, finded.Conditions[0])
		require.Equal(t, updatedCondition, finded.Conditions[0])
	})

	t.Run("FindByUserId after DeleteCondition", func(t *testing.T) {
		conditionRepo := initConditionRepo(t)
		ctx := context.TODO()

		success, err := conditionRepo.InitDesiredCondition(ctx, userId)
		require.NoError(t, err)
		require.True(t, success)

		savedCondition := newCondition()
		success, err = conditionRepo.InsertCondition(ctx, userId, 1, savedCondition)
		require.NoError(t, err)
		require.True(t, success)

		success, err = conditionRepo.DeleteCondition(ctx, userId, savedCondition.ConditionId)
		require.NoError(t, err)
		require.True(t, success)

		finded, err := conditionRepo.FindByUserId(ctx, userId)
		require.NoError(t, err)
		require.Len(t, finded.Conditions, 0)
	})

	limitConditions := []*condition.Condition{
		{
			ConditionId:   "conditionId1",
			ConditionName: "test_condition_name",
			Query:         &condition.Query{},
		},
		{
			ConditionId:   "conditionId2",
			ConditionName: "test_condition_name",
			Query:         &condition.Query{},
		},
		{
			ConditionId:   "conditionId3",
			ConditionName: "test_condition_name",
			Query:         &condition.Query{},
		},
		{
			ConditionId:   "conditionId4",
			ConditionName: "test_condition_name",
			Query:         &condition.Query{},
		},
	}
	t.Run("InsertCondition with limitCount", func(t *testing.T) {
		conditionRepo := initConditionRepo(t)
		ctx := context.TODO()

		success, err := conditionRepo.InitDesiredCondition(ctx, userId)
		require.NoError(t, err)
		require.True(t, success)

		limitCount := uint(len(limitConditions) - 1)
		for _, c := range limitConditions[0:limitCount] {
			success, err := conditionRepo.InsertCondition(ctx, userId, limitCount, c)
			require.NoError(t, err)
			require.True(t, success)
		}

		success, err = conditionRepo.InsertCondition(ctx, userId, limitCount, limitConditions[limitCount])
		require.NoError(t, err)
		require.False(t, success)
	})

	t.Run("unique userId", func(t *testing.T) {
		conditionRepo := initConditionRepo(t)
		ctx := context.TODO()

		success, err := conditionRepo.InitDesiredCondition(ctx, userId)
		require.NoError(t, err)
		require.True(t, success)
		success, err = conditionRepo.InitDesiredCondition(ctx, userId)
		require.NoError(t, err)
		require.False(t, success)
	})

	t.Run("set conditionId when insert condition", func(t *testing.T) {
		conditionRepo := initConditionRepo(t)
		ctx := context.TODO()

		success, err := conditionRepo.InitDesiredCondition(ctx, userId)
		require.NoError(t, err)
		require.True(t, success)

		savedCondition := newCondition()
		success, err = conditionRepo.InsertCondition(ctx, userId, 2, savedCondition)
		firstConditionId := savedCondition.ConditionId
		require.NoError(t, err)
		require.True(t, success)
		require.NotEmpty(t, firstConditionId)

		success, err = conditionRepo.InsertCondition(ctx, userId, 2, savedCondition)
		secondConditionId := savedCondition.ConditionId
		require.NoError(t, err)
		require.True(t, success)
		require.NotEmpty(t, secondConditionId)

		require.NotEqual(t, firstConditionId, secondConditionId)
	})

	t.Run("limitCount can't be zero", func(t *testing.T) {
		conditionRepo := initConditionRepo(t)
		ctx := context.TODO()

		success, err := conditionRepo.InitDesiredCondition(ctx, userId)
		require.NoError(t, err)
		require.True(t, success)

		success, err = conditionRepo.InsertCondition(ctx, userId, 0, newCondition())
		require.Error(t, err)
		require.Equal(t, repo.ErrNonZero, err)
		require.False(t, success)
	})
}

func initConditionRepo(t *testing.T) repo.ConditionRepo {
	db := tinit.InitDB(t)
	return tinit.InitRestapiConditionRepo(t, db)
}

func checkSimilarTimes(t *testing.T, after time.Time, before time.Time) {
	require.GreaterOrEqual(t, after.UTC(), before.UTC())
	require.LessOrEqual(t, after.UTC().Add(time.Second*-1), before.UTC())
}

func newCondition() *condition.Condition {
	return &condition.Condition{
		ConditionName: "test_condition_name",
		Query: &condition.Query{
			Categories: []*condition.CategoryQuery{
				{
					Site:         "test_site",
					CategoryName: "test_category_name",
				},
			},
			SkillNames: [][]string{{"test_skill_name"}},
			MinCareer:  ptr.P(int32(1)),
			MaxCareer:  ptr.P(int32(2)),
		},
	}
}

func newUpdatedCondition(conditionId string) *condition.Condition {
	return &condition.Condition{
		ConditionId:   conditionId,
		ConditionName: "update_condition_name",
		Query: &condition.Query{
			Categories: []*condition.CategoryQuery{
				{
					Site:         "update_site",
					CategoryName: "update_category_name",
				},
			},
			SkillNames: [][]string{{"update_skill_name"}},
			MinCareer:  ptr.P(int32(3)),
			MaxCareer:  nil,
		},
	}
}
