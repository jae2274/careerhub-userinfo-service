package service

import (
	"context"
	"time"

	"github.com/jae2274/careerhub-userinfo-service/careerhub/userinfo_service/common/domain/condition"
	"github.com/jae2274/careerhub-userinfo-service/careerhub/userinfo_service/common/domain/suggestion"
	"github.com/jae2274/careerhub-userinfo-service/careerhub/userinfo_service/suggester/repo"
)

type SuggesterService struct {
	conditionRepo  repo.ConditionRepo
	historyRepo    repo.HistoryRepo
	suggestionRepo repo.SuggestionRepo
}

func NewSuggesterService(conditionRepo repo.ConditionRepo, historyRepo repo.HistoryRepo, suggestionRepo repo.SuggestionRepo) *SuggesterService {
	return &SuggesterService{
		conditionRepo:  conditionRepo,
		historyRepo:    historyRepo,
		suggestionRepo: suggestionRepo,
	}
}

func (s *SuggesterService) FindWorkingBatchIds(ctx context.Context) ([]string, error) {
	return s.historyRepo.FindWorkingBatchIds(ctx)
}

func (s *SuggesterService) UpdateFailed(ctx context.Context, batchIds []string) error {
	if len(batchIds) == 0 {
		return nil
	}

	err := s.suggestionRepo.DeleteSuggestions(ctx, batchIds)
	if err != nil {
		return err
	}
	err = s.historyRepo.UpdateFailed(ctx, batchIds)
	if err != nil {
		return err
	}

	return nil
}

func (s *SuggesterService) StartBatch(ctx context.Context, batchId string) (*time.Time, error) {
	lastSuccessedDate, err := s.historyRepo.FindLastSuccessedDate(ctx)
	if err != nil {
		return nil, err
	}

	err = s.historyRepo.InsertHistory(ctx, batchId)
	if err != nil {
		return nil, err
	}

	return lastSuccessedDate, nil
}

func (s *SuggesterService) GetDesiredConditions(ctx context.Context) ([]condition.DesiredCondition, error) {
	return s.conditionRepo.GetDesiredConditions(ctx)
}

func (s *SuggesterService) InsertSuggestion(ctx context.Context, suggestion *suggestion.Suggestion) error {
	return s.suggestionRepo.InsertSuggestion(ctx, suggestion)
}

func (s *SuggesterService) EndBatch(ctx context.Context, batchId string) error {
	return s.historyRepo.UpdateSuccessed(ctx, batchId)
}
