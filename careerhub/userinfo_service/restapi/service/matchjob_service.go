package service

import (
	"context"

	condition "github.com/jae2274/careerhub-userinfo-service/careerhub/userinfo_service/common/domain/matchjob"
	"github.com/jae2274/careerhub-userinfo-service/careerhub/userinfo_service/restapi/repo"
)

type MatchJobService interface {
	FindByUserId(ctx context.Context, userId string) (*condition.MatchJob, error)
	InsertCondition(ctx context.Context, userId string, limitCount uint, newCondition *condition.Condition) (bool, error)
	UpdateCondition(ctx context.Context, userId string, updateCondition *condition.Condition) (bool, error)
	DeleteCondition(ctx context.Context, userId string, conditionId string) (bool, error)
	UpdateAgreeToMail(ctx context.Context, userId string, agreeToMail bool) (bool, error)
}

type MatchJobServiceImpl struct {
	matchJobRepo repo.MatchJobRepo
}

func NewMatchJobService(matchJobRepo repo.MatchJobRepo) MatchJobService {
	return &MatchJobServiceImpl{
		matchJobRepo: matchJobRepo,
	}
}

func (c *MatchJobServiceImpl) FindByUserId(ctx context.Context, userId string) (*condition.MatchJob, error) {
	matchJob, err := c.matchJobRepo.FindByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}

	if matchJob == nil {
		_, err := c.matchJobRepo.InitMatchJob(ctx, userId)
		if err != nil {
			return nil, err
		}

		matchJob, err = c.matchJobRepo.FindByUserId(ctx, userId)
		if err != nil {
			return nil, err
		}
	}

	return matchJob, nil
}

func (c *MatchJobServiceImpl) InsertCondition(ctx context.Context, userId string, limitCount uint, newCondition *condition.Condition) (bool, error) {
	_, err := c.matchJobRepo.InitMatchJob(ctx, userId)
	if err != nil {
		return false, err
	}

	return c.matchJobRepo.InsertCondition(ctx, userId, limitCount, newCondition)
}

func (c *MatchJobServiceImpl) UpdateCondition(ctx context.Context, userId string, updateCondition *condition.Condition) (bool, error) {
	_, err := c.matchJobRepo.InitMatchJob(ctx, userId)
	if err != nil {
		return false, err
	}

	return c.matchJobRepo.UpdateCondition(ctx, userId, updateCondition)
}

func (c *MatchJobServiceImpl) DeleteCondition(ctx context.Context, userId string, conditionId string) (bool, error) {
	_, err := c.matchJobRepo.InitMatchJob(ctx, userId)
	if err != nil {
		return false, err
	}

	return c.matchJobRepo.DeleteCondition(ctx, userId, conditionId)
}

func (c *MatchJobServiceImpl) UpdateAgreeToMail(ctx context.Context, userId string, agreeToMail bool) (bool, error) {
	_, err := c.matchJobRepo.InitMatchJob(ctx, userId)
	if err != nil {
		return false, err
	}

	return c.matchJobRepo.UpdateAgreeToMail(ctx, userId, agreeToMail)
}
