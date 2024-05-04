package restapi

import (
	"context"
	"fmt"
	"net"

	"github.com/jae2274/careerhub-userinfo-service/careerhub/userinfo_service/restapi/repo"
	"github.com/jae2274/careerhub-userinfo-service/careerhub/userinfo_service/restapi/restapi_grpc"
	"github.com/jae2274/careerhub-userinfo-service/careerhub/userinfo_service/restapi/server"
	"github.com/jae2274/careerhub-userinfo-service/careerhub/userinfo_service/restapi/service"
	"github.com/jae2274/careerhub-userinfo-service/careerhub/userinfo_service/utils"
	"github.com/jae2274/goutils/llog"
	"github.com/jae2274/goutils/terr"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

func Run(ctx context.Context, grpcPort int, db *mongo.Database) error {
	matchJobGrpcServer := runMatchJobGrpcServer(ctx, grpcPort, db)
	scrapJobGrpcServer := runScrapJobGrpcServer(ctx, grpcPort, db)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		return terr.Wrap(err)
	}

	llog.Msg("Start restapi grpc server").Level(llog.INFO).Data("port", grpcPort).Log(ctx)

	grpcServer := grpc.NewServer(utils.Middlewares()...)
	restapi_grpc.RegisterMatchJobGrpcServer(grpcServer, matchJobGrpcServer)
	restapi_grpc.RegisterScrapJobGrpcServer(grpcServer, scrapJobGrpcServer)

	err = grpcServer.Serve(listener)
	if err != nil {
		return terr.Wrap(err)
	}

	return nil
}

func runMatchJobGrpcServer(ctx context.Context, grpcPort int, db *mongo.Database) restapi_grpc.MatchJobGrpcServer {
	matchJobRepo := repo.NewMatchJobRepo(db)
	conditionService := service.NewMatchJobService(matchJobRepo)

	return server.NewMatchJobGrpcServer(conditionService)
}

func runScrapJobGrpcServer(ctx context.Context, grpcPort int, db *mongo.Database) restapi_grpc.ScrapJobGrpcServer {
	scrapJobRepo := repo.NewScrapJobRepo(db)

	return server.NewScrapJobGrpcServer(scrapJobRepo)
}
