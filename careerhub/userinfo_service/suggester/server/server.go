package server

import (
	"context"

	"github.com/jae2274/careerhub-userinfo-service/careerhub/userinfo_service/common/domain/condition"
	"github.com/jae2274/careerhub-userinfo-service/careerhub/userinfo_service/suggester/repo"
	"github.com/jae2274/careerhub-userinfo-service/careerhub/userinfo_service/suggester/suggester_grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type SuggesterGrpcServer struct {
	conditionRepo repo.ConditionRepo
	suggester_grpc.UnimplementedUserinfoServer
}

func NewSuggesterGrpcServer(conditionRepo repo.ConditionRepo) *SuggesterGrpcServer {
	return &SuggesterGrpcServer{
		conditionRepo: conditionRepo,
	}
}
func (s *SuggesterGrpcServer) GetConditions(ctx context.Context, _ *emptypb.Empty) (*suggester_grpc.GetConditionsResponse, error) {
	desiredConditions, err := s.conditionRepo.GetDesiredConditions(ctx)
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
