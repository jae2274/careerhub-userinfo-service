package server

import (
	"context"

	"github.com/jae2274/careerhub-userinfo-service/careerhub/userinfo_service/common/domain/scrapjob"
	"github.com/jae2274/careerhub-userinfo-service/careerhub/userinfo_service/restapi/repo"
	"github.com/jae2274/careerhub-userinfo-service/careerhub/userinfo_service/restapi/restapi_grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ScrapJobGrpcServer struct {
	scrapJobRepo repo.ScrapJobRepo
	restapi_grpc.UnimplementedScrapJobGrpcServer
}

func NewScrapJobGrpcServer(scrapJobRepo repo.ScrapJobRepo) restapi_grpc.ScrapJobGrpcServer {
	return &ScrapJobGrpcServer{
		scrapJobRepo: scrapJobRepo,
	}
}

func (s *ScrapJobGrpcServer) GetScrapJobs(ctx context.Context, in *restapi_grpc.GetScrapJobsRequest) (*restapi_grpc.GetScrapJobsResponse, error) {
	scrapJobs, err := s.scrapJobRepo.GetScrapJobs(ctx, in.UserId, in.Tag)
	if err != nil {
		return nil, err
	}

	grpcScrapJobs := convertScrapJobsToGrpc(scrapJobs)

	return &restapi_grpc.GetScrapJobsResponse{ScrapJobs: grpcScrapJobs}, nil
}

func (s *ScrapJobGrpcServer) AddScrapJob(ctx context.Context, in *restapi_grpc.AddScrapJobRequest) (*emptypb.Empty, error) {
	scrapJob := &scrapjob.ScrapJob{
		UserId:    in.UserId,
		Site:      in.Site,
		PostingId: in.PostingId,
		Tags:      []string{},
	}

	err := s.scrapJobRepo.AddScrapJob(ctx, scrapJob)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *ScrapJobGrpcServer) RemoveScrapJob(ctx context.Context, in *restapi_grpc.RemoveScrapJobRequest) (*restapi_grpc.IsExistedResponse, error) {
	isExisted, err := s.scrapJobRepo.RemoveScrapJob(ctx, in.UserId, in.Site, in.PostingId)
	if err != nil {
		return nil, err
	}

	return &restapi_grpc.IsExistedResponse{IsExisted: isExisted}, nil
}

func (s *ScrapJobGrpcServer) AddTag(ctx context.Context, in *restapi_grpc.AddTagRequest) (*restapi_grpc.IsExistedResponse, error) {
	isExisted, err := s.scrapJobRepo.AddTag(ctx, in.UserId, in.Site, in.PostingId, in.Tag)
	if err != nil {
		return nil, err
	}

	return &restapi_grpc.IsExistedResponse{IsExisted: isExisted}, nil
}

func (s *ScrapJobGrpcServer) RemoveTag(ctx context.Context, in *restapi_grpc.RemoveTagRequest) (*restapi_grpc.IsExistedResponse, error) {
	isExisted, err := s.scrapJobRepo.RemoveTag(ctx, in.UserId, in.Site, in.PostingId, in.Tag)
	if err != nil {
		return nil, err
	}

	return &restapi_grpc.IsExistedResponse{IsExisted: isExisted}, nil
}

func (s *ScrapJobGrpcServer) GetScrapTags(ctx context.Context, in *restapi_grpc.GetScrapTagsRequest) (*restapi_grpc.GetScrapTagsResponse, error) {
	tags, err := s.scrapJobRepo.GetScrapTags(ctx, in.UserId)
	if err != nil {
		return nil, err
	}

	return &restapi_grpc.GetScrapTagsResponse{Tags: tags}, nil
}

func (s *ScrapJobGrpcServer) GetScrapJobsById(ctx context.Context, in *restapi_grpc.GetScrapJobsByIdRequest) (*restapi_grpc.GetScrapJobsResponse, error) {
	scrapJobs, err := s.scrapJobRepo.GetScrapJobsById(ctx, in.UserId, in.JobPostingIds)
	if err != nil {
		return nil, err
	}

	grpcScrapJobs := convertScrapJobsToGrpc(scrapJobs)

	return &restapi_grpc.GetScrapJobsResponse{ScrapJobs: grpcScrapJobs}, nil
}

func (s *ScrapJobGrpcServer) GetUntaggedScrapJobs(ctx context.Context, in *restapi_grpc.GetUntaggedScrapJobsRequest) (*restapi_grpc.GetScrapJobsResponse, error) {
	scrapJobs, err := s.scrapJobRepo.GetUntaggedScrapJobs(ctx, in.UserId)
	if err != nil {
		return nil, err
	}

	grpcScrapJobs := convertScrapJobsToGrpc(scrapJobs)

	return &restapi_grpc.GetScrapJobsResponse{ScrapJobs: grpcScrapJobs}, nil
}

func convertScrapJobsToGrpc(scrapJobs []*scrapjob.ScrapJob) []*restapi_grpc.ScrapJob {
	grpcScrapJobs := make([]*restapi_grpc.ScrapJob, len(scrapJobs))
	for i, scrapJob := range scrapJobs {
		grpcScrapJobs[i] = &restapi_grpc.ScrapJob{
			Site:      scrapJob.Site,
			PostingId: scrapJob.PostingId,
			Tags:      scrapJob.Tags,
		}
	}
	return grpcScrapJobs
}
