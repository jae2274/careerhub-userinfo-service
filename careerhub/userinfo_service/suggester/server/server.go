package server

import (
	"context"

	condition "github.com/jae2274/careerhub-userinfo-service/careerhub/userinfo_service/common/domain/matchjob"
	"github.com/jae2274/careerhub-userinfo-service/careerhub/userinfo_service/suggester/repo"
	"github.com/jae2274/careerhub-userinfo-service/careerhub/userinfo_service/suggester/suggester_grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type SuggesterGrpcServer struct {
	matchJobRepo repo.MatchJobRepo
	suggester_grpc.UnimplementedUserinfoServer
}

func NewSuggesterGrpcServer(matchJobRepo repo.MatchJobRepo) *SuggesterGrpcServer {
	return &SuggesterGrpcServer{
		matchJobRepo: matchJobRepo,
	}
}
func (s *SuggesterGrpcServer) GetConditions(ctx context.Context, _ *emptypb.Empty) (*suggester_grpc.GetConditionsResponse, error) {
	matchJobs, err := s.matchJobRepo.GetMatchJobs(ctx)
	if err != nil {
		return nil, err
	}

	return &suggester_grpc.GetConditionsResponse{
		Conditions: convertMatchJobsToGrpc(matchJobs),
	}, nil
}

func convertMatchJobsToGrpc(matchJobs []*condition.MatchJob) []*suggester_grpc.Condition {
	var grpcConditions []*suggester_grpc.Condition

	for _, matchJob := range matchJobs {
		for _, c := range matchJob.Conditions {
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
				UserId:        matchJob.UserId,
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
