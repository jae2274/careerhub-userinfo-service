package server

import (
	"context"
	"time"

	"github.com/jae2274/careerhub-userinfo-service/careerhub/userinfo_service/common/domain/condition"
	"github.com/jae2274/careerhub-userinfo-service/careerhub/userinfo_service/common/domain/suggestion"
	"github.com/jae2274/careerhub-userinfo-service/careerhub/userinfo_service/suggester/service"
	"github.com/jae2274/careerhub-userinfo-service/careerhub/userinfo_service/suggester/suggester_grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type SuggesterGrpcServer struct {
	service *service.SuggesterService
	suggester_grpc.UnimplementedSuggesterGrpcServer
}

func NewSuggesterGrpcServer(service *service.SuggesterService) *SuggesterGrpcServer {
	return &SuggesterGrpcServer{
		service: service,
	}
}

func (s *SuggesterGrpcServer) StartBatch(ctx context.Context, req *suggester_grpc.StartBatchRequest) (*suggester_grpc.StartBatchResponse, error) {
	batchIds, err := s.service.FindWorkingBatchIds(ctx)
	if err != nil {
		return nil, err
	}
	s.service.UpdateFailed(ctx, batchIds)

	lastWorkedDate, err := s.service.StartBatch(ctx, req.BatchId, time.UnixMilli(req.StartTimeUnixMilli))
	if err != nil {
		return nil, err
	}

	if lastWorkedDate == nil {
		yesterday := time.Now().Add(-24 * time.Hour)
		lastWorkedDate = &yesterday
	}

	return &suggester_grpc.StartBatchResponse{
		AfterUnixMilli: lastWorkedDate.UnixMilli(),
	}, nil
}

func (s *SuggesterGrpcServer) GetConditions(ctx context.Context, _ *emptypb.Empty) (*suggester_grpc.GetConditionsResponse, error) {
	desiredConditions, err := s.service.GetDesiredConditions(ctx)
	if err != nil {
		return nil, err
	}

	return &suggester_grpc.GetConditionsResponse{
		Conditions: convertDesiredConditionsToGrpc(desiredConditions),
	}, nil
}

func convertDesiredConditionsToGrpc(desiredConditions []*condition.DesiredCondition) []*suggester_grpc.Condition {
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
	return &emptypb.Empty{}, s.service.InsertSuggestion(ctx, convertSuggestionToDomain(req))
}

func convertSuggestionToDomain(sg *suggester_grpc.Suggestion) *suggestion.Suggestion {
	postings := make([]*suggestion.Posting, len(sg.Postings))
	for i, posting := range sg.Postings {
		postings[i] = &suggestion.Posting{
			PostingId: &suggestion.PostingId{
				Site:      posting.PostingId.Site,
				PostingId: posting.PostingId.PostingId,
			},
			Title:       posting.Title,
			CompanyId:   posting.CompanyId,
			CompanyName: posting.CompanyName,
			Info: &suggestion.PostingInfo{
				Categories: posting.PostingInfo.Categories,
				SkillNames: posting.PostingInfo.SkillNames,
				MinCareer:  posting.PostingInfo.MinCareer,
				MaxCareer:  posting.PostingInfo.MaxCareer,
			},
		}
	}

	return &suggestion.Suggestion{
		BatchId:       sg.BatchId,
		UserId:        sg.UserId,
		ConditionId:   sg.ConditionId,
		ConditionName: sg.ConditionName,
		Postings:      postings,
	}
}

func (s *SuggesterGrpcServer) EndBatch(ctx context.Context, req *suggester_grpc.EndBatchRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, s.service.EndBatch(ctx, req.BatchId, time.UnixMilli(req.EndTimeUnixMilli))
}
