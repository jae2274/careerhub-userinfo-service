package server

import (
	"context"

	"github.com/jae2274/careerhub-userinfo-service/careerhub/userinfo_service/common/domain/condition"
	"github.com/jae2274/careerhub-userinfo-service/careerhub/userinfo_service/suggester/repo"
	"github.com/jae2274/careerhub-userinfo-service/careerhub/userinfo_service/suggester/suggester_grpc"
	"github.com/jae2274/goutils/terr"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type SuggesterGrpcServer struct {
	conditionRepo repo.ConditionRepo
	historyRepo   repo.HistoryRepo
	suggester_grpc.UnimplementedSuggesterGrpcServer
}

func NewSuggesterGrpcServer(conditionRepo repo.ConditionRepo, historyRepo repo.HistoryRepo) *SuggesterGrpcServer {
	return &SuggesterGrpcServer{
		conditionRepo: conditionRepo,
		historyRepo:   historyRepo,
	}
}

func (s *SuggesterGrpcServer) StartBatch(ctx context.Context, req *suggester_grpc.StartBatchRequest) (*suggester_grpc.StartBatchResponse, error) {
	lastWorkedDate, err := s.historyRepo.GetLastSuccessDate(ctx)
	if err != nil {
		return nil, terr.Wrap(err)
	}

	err = s.historyRepo.InsertHistory(ctx, req.BatchId)
	if err != nil {
		return nil, terr.Wrap(err)
	}

	return &suggester_grpc.StartBatchResponse{
		AfterUnixMilli: lastWorkedDate.UnixMilli(),
	}, nil
}

func (s *SuggesterGrpcServer) GetConditions(ctx context.Context, _ *emptypb.Empty) (*suggester_grpc.GetConditionsResponse, error) {
	desiredConditions, err := s.conditionRepo.GetConditions(ctx)
	if err != nil {
		return nil, terr.Wrap(err)
	}

	return &suggester_grpc.GetConditionsResponse{
		Conditions: convertDesiredConditionsToGrpc(desiredConditions),
	}, nil
}

func convertDesiredConditionsToGrpc(desiredConditions []condition.DesiredCondition) []*suggester_grpc.Condition {
	var grpcConditions []*suggester_grpc.Condition

	for _, desiredCondition := range desiredConditions {
		for _, c := range desiredCondition.Conditions {
			categories := make([]*suggester_grpc.Category, len(c.Query.Categories))
			for i, category := range c.Query.Categories {
				categories[i] = &suggester_grpc.Category{
					Site:         category.Site,
					CategoryName: category.CategoryName,
				}
			}

			skillNames := make([]*suggester_grpc.Skill, len(c.Query.SkillNames))
			for i, skills := range c.Query.SkillNames {
				skillNames[i] = &suggester_grpc.Skill{
					Or: skills,
				}
			}

			grpcConditions = append(grpcConditions, &suggester_grpc.Condition{
				UserId:        desiredCondition.UserId,
				ConditionId:   c.ConditionId,
				ConditionName: c.ConditionName,
				Query: &suggester_grpc.Query{
					Categories: categories,
					SkillNames: skillNames,
					MinCareer:  c.Query.MinCareer,
					MaxCareer:  c.Query.MaxCareer,
				},
			})
		}
	}

	return grpcConditions
}

func (s *SuggesterGrpcServer) ReceiveSuggestion(ctx context.Context, req *suggester_grpc.Suggestion) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReceiveSuggestions not implemented")
}

func (s *SuggesterGrpcServer) EndBatch(ctx context.Context, req *suggester_grpc.EndBatchRequest) (*emptypb.Empty, error) {
	err := s.historyRepo.UpdateHistory(ctx, req.BatchId)

	if err != nil {
		return nil, terr.Wrap(err)
	}

	return &emptypb.Empty{}, nil
}
