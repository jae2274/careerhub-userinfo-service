package service

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/jae2274/careerhub-userinfo-service/careerhub/userinfo_service/common/domain/condition"
	"github.com/jae2274/careerhub-userinfo-service/careerhub/userinfo_service/restapi/repo"
	"github.com/jae2274/careerhub-userinfo-service/careerhub/userinfo_service/restapi/service"
	"github.com/jae2274/careerhub-userinfo-service/test/tinit"
	"github.com/jae2274/goutils/ptr"
	"github.com/stretchr/testify/require"
)

func TestConditions(t *testing.T) {
	t.Run("return empty", func(t *testing.T) {
		ctx := context.Background()
		svc := initService(t)

		desiredCondition, err := svc.FindByUserId(ctx, "userId")
		require.NoError(t, err)

		require.Equal(t, "userId", desiredCondition.UserId)
		require.False(t, desiredCondition.AgreeToMail)
		require.Len(t, desiredCondition.Conditions, 0)
	})

	t.Run("return one condition after insert", func(t *testing.T) {
		ctx := context.Background()
		svc := initService(t)

		savedCondition := newCondition()

		isSuccess, err := svc.InsertCondition(ctx, "userId", 1, savedCondition)
		require.NoError(t, err)
		require.True(t, isSuccess)

		desiredCondition, err := svc.FindByUserId(ctx, "userId")
		require.NoError(t, err)

		require.Equal(t, "userId", desiredCondition.UserId)
		require.False(t, desiredCondition.AgreeToMail)
		require.Len(t, desiredCondition.Conditions, 1)

		require.Equal(t, savedCondition, desiredCondition.Conditions[0])
	})

	t.Run("return updated condition after update", func(t *testing.T) {
		ctx := context.Background()
		svc := initService(t)

		savedCondition := newCondition()

		isSuccess, err := svc.InsertCondition(ctx, "userId", 1, savedCondition)
		require.NoError(t, err)
		require.True(t, isSuccess)

		updatedCondition := newUpdatedCondition(savedCondition.ConditionId)
		isSuccess, err = svc.UpdateCondition(ctx, "userId", updatedCondition)
		require.NoError(t, err)
		require.True(t, isSuccess)

		desiredCondition, err := svc.FindByUserId(ctx, "userId")
		require.NoError(t, err)

		require.Equal(t, "userId", desiredCondition.UserId)
		require.False(t, desiredCondition.AgreeToMail)
		require.Len(t, desiredCondition.Conditions, 1)

		require.Equal(t, updatedCondition, desiredCondition.Conditions[0])
	})

	t.Run("return empty after delete", func(t *testing.T) {
		ctx := context.Background()
		svc := initService(t)

		savedCondition := newCondition()

		isSuccess, err := svc.InsertCondition(ctx, "userId", 1, savedCondition)
		require.NoError(t, err)
		require.True(t, isSuccess)

		isSuccess, err = svc.DeleteCondition(ctx, "userId", savedCondition.ConditionId)
		require.NoError(t, err)
		require.True(t, isSuccess)

		desiredCondition, err := svc.FindByUserId(ctx, "userId")
		require.NoError(t, err)

		require.Equal(t, "userId", desiredCondition.UserId)
		require.False(t, desiredCondition.AgreeToMail)
		require.Len(t, desiredCondition.Conditions, 0)
	})

	t.Run("return error when insert condition with zero limit", func(t *testing.T) {
		ctx := context.Background()
		svc := initService(t)

		savedCondition := newCondition()

		isSuccess, err := svc.InsertCondition(ctx, "userId", 0, savedCondition)
		require.Error(t, err)
		require.False(t, isSuccess)
	})

	t.Run("return false when insert condition over limit", func(t *testing.T) {
		ctx := context.Background()
		svc := initService(t)

		limitCount := uint(3)
		for i := 0; i < 3; i++ {
			isSuccess, err := svc.InsertCondition(ctx, "userId", limitCount, newCondition())
			require.NoError(t, err, i)
			require.True(t, isSuccess, i)
		}

		isSuccess, err := svc.InsertCondition(ctx, "userId", limitCount, newCondition())
		require.NoError(t, err)
		require.False(t, isSuccess)
	})

	t.Run("set conditionId when insert condition", func(t *testing.T) {
		ctx := context.Background()
		svc := initService(t)

		savedCondition := newCondition()

		isSuccess, err := svc.InsertCondition(ctx, "userId", 3, savedCondition)
		firstConditionId := savedCondition.ConditionId
		require.NoError(t, err)
		require.True(t, isSuccess)
		require.NotEmpty(t, firstConditionId)

		isSuccess, err = svc.InsertCondition(ctx, "userId", 3, savedCondition)
		secondConditionId := savedCondition.ConditionId
		require.NoError(t, err)
		require.True(t, isSuccess)
		require.NotEmpty(t, secondConditionId)

		require.NotEqual(t, firstConditionId, secondConditionId)
	})

	t.Run("return false when update condition with non-exist conditionId", func(t *testing.T) {
		ctx := context.Background()
		svc := initService(t)

		isSuccess, err := svc.UpdateCondition(ctx, "userId", newUpdatedCondition(uuid.NewString()))
		require.NoError(t, err)
		require.False(t, isSuccess)
	})

	t.Run("return false when delete condition with non-exist conditionId", func(t *testing.T) {
		ctx := context.Background()
		svc := initService(t)

		isSuccess, err := svc.DeleteCondition(ctx, "userId", uuid.NewString())
		require.NoError(t, err)
		require.False(t, isSuccess)
	})
}

func TestAgreeToMail(t *testing.T) {

	testUpdateAgreeToMail := func(t *testing.T, agreeToMail bool) {
		ctx := context.Background()
		svc := initService(t)

		userId := "userId"
		isSuccess, err := svc.UpdateAgreeToMail(ctx, userId, agreeToMail)
		require.NoError(t, err)
		require.True(t, isSuccess)

		desiredCondition, err := svc.FindByUserId(ctx, userId)
		require.NoError(t, err)

		require.Equal(t, userId, desiredCondition.UserId)
		require.Equal(t, agreeToMail, desiredCondition.AgreeToMail)
	}
	t.Run("after update agreeToMail to false", func(t *testing.T) {
		testUpdateAgreeToMail(t, false)
	})

	t.Run("after update agreeToMail to true", func(t *testing.T) {
		testUpdateAgreeToMail(t, true)
	})
}

func initService(t *testing.T) service.ConditionService {
	conditionRepo := repo.NewConditionRepo(tinit.InitDB(t))

	return service.NewConditionService(conditionRepo)
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
