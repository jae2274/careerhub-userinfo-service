package repo

import (
	"context"
	"testing"
	"time"

	"github.com/jae2274/careerhub-userinfo-service/careerhub/userinfo_service/restapi/repo"
	"github.com/jae2274/careerhub-userinfo-service/test/tinit"
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

		err := conditionRepo.InitConditions(ctx, userId)
		require.NoError(t, err)

		desiredCondition, err := conditionRepo.FindByUserId(ctx, userId)
		require.NoError(t, err)
		require.NotNil(t, desiredCondition)
		require.Equal(t, userId, desiredCondition.UserId)

		now := time.Now()
		checkSimilarTimes(t, now, desiredCondition.InsertedAt)
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
