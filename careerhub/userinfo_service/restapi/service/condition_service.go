package service

import (
	"context"

	"github.com/jae2274/careerhub-userinfo-service/careerhub/userinfo_service/common/domain/condition"
	"github.com/jae2274/careerhub-userinfo-service/careerhub/userinfo_service/restapi/repo"
)

type ConditionService interface {
	FindByUserId(ctx context.Context, userId string) (*condition.DesiredCondition, error)
	InsertCondition(ctx context.Context, userId string, limitCount uint, newCondition *condition.Condition) (bool, error)
	UpdateCondition(ctx context.Context, userId string, updateCondition *condition.Condition) (bool, error)
	DeleteCondition(ctx context.Context, userId string, conditionId string) (bool, error)
	UpdateAgreeToMail(ctx context.Context, userId string, agreeToMail bool) (bool, error)
}

type ConditionServiceImpl struct {
	conditionRepo repo.ConditionRepo
}

func NewConditionService(conditionRepo repo.ConditionRepo) ConditionService {
	return &ConditionServiceImpl{
		conditionRepo: conditionRepo,
	}
}

func (c *ConditionServiceImpl) FindByUserId(ctx context.Context, userId string) (*condition.DesiredCondition, error) {
	desiredCondition, err := c.conditionRepo.FindByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}

	if desiredCondition == nil {
		_, err := c.conditionRepo.InitDesiredCondition(ctx, userId)
		if err != nil {
			return nil, err
		}

		desiredCondition, err = c.conditionRepo.FindByUserId(ctx, userId)
		if err != nil {
			return nil, err
		}
	}

	return desiredCondition, nil
}

func (c *ConditionServiceImpl) InsertCondition(ctx context.Context, userId string, limitCount uint, newCondition *condition.Condition) (bool, error) {
	_, err := c.conditionRepo.InitDesiredCondition(ctx, userId)
	if err != nil {
		return false, err
	}

	return c.conditionRepo.InsertCondition(ctx, userId, limitCount, newCondition)
}

func (c *ConditionServiceImpl) UpdateCondition(ctx context.Context, userId string, updateCondition *condition.Condition) (bool, error) {
	_, err := c.conditionRepo.InitDesiredCondition(ctx, userId)
	if err != nil {
		return false, err
	}

	return c.conditionRepo.UpdateCondition(ctx, userId, updateCondition)
}

func (c *ConditionServiceImpl) DeleteCondition(ctx context.Context, userId string, conditionId string) (bool, error) {
	_, err := c.conditionRepo.InitDesiredCondition(ctx, userId)
	if err != nil {
		return false, err
	}

	return c.conditionRepo.DeleteCondition(ctx, userId, conditionId)
}

func (c *ConditionServiceImpl) UpdateAgreeToMail(ctx context.Context, userId string, agreeToMail bool) (bool, error) {
	_, err := c.conditionRepo.InitDesiredCondition(ctx, userId)
	if err != nil {
		return false, err
	}

	return c.conditionRepo.UpdateAgreeToMail(ctx, userId, agreeToMail)
}
