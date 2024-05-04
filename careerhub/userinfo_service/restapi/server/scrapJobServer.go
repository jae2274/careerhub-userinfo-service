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
	scrapJobs, err := s.scrapJobRepo.GetScrapJobs(ctx, in.UserId)
	if err != nil {
		return nil, err
	}

	grpcScrapJobs := make([]*restapi_grpc.ScrapJob, len(scrapJobs))
	for i, scrapJob := range scrapJobs {
		grpcScrapJobs[i] = &restapi_grpc.ScrapJob{
			Site:      scrapJob.Site,
			PostingId: scrapJob.PostingId,
		}
	}

	return &restapi_grpc.GetScrapJobsResponse{ScrapJobs: grpcScrapJobs}, nil
}

func (s *ScrapJobGrpcServer) AddScrapJob(ctx context.Context, in *restapi_grpc.AddScrapJobRequest) (*emptypb.Empty, error) {
	scrapJob := &scrapjob.ScrapJob{
		UserId:    in.UserId,
		Site:      in.Site,
		PostingId: in.PostingId,
	}

	err := s.scrapJobRepo.AddScrapJob(ctx, scrapJob)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *ScrapJobGrpcServer) RemoveScrapJob(ctx context.Context, in *restapi_grpc.RemoveScrapJobRequest) (*emptypb.Empty, error) {
	err := s.scrapJobRepo.RemoveScrapJob(ctx, in.UserId, in.Site, in.PostingId)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
