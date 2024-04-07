package suggester

import (
	"context"
	"fmt"
	"net"

	"github.com/jae2274/careerhub-userinfo-service/careerhub/userinfo_service/suggester/repo"
	"github.com/jae2274/careerhub-userinfo-service/careerhub/userinfo_service/suggester/server"
	"github.com/jae2274/careerhub-userinfo-service/careerhub/userinfo_service/suggester/service"
	"github.com/jae2274/careerhub-userinfo-service/careerhub/userinfo_service/suggester/suggester_grpc"
	"github.com/jae2274/careerhub-userinfo-service/careerhub/userinfo_service/utils"
	"github.com/jae2274/goutils/llog"
	"github.com/jae2274/goutils/terr"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

func Run(ctx context.Context, grpcPort int, db *mongo.Database) error {
	conditionRepo := repo.NewConditionRepo(db)
	suggestionRepo := repo.NewSuggestionRepo(db)
	historyRepo := repo.NewHistoryRepo(db)
	service := service.NewSuggesterService(conditionRepo, historyRepo, suggestionRepo)
	suggesterGrpcserver := server.NewSuggesterGrpcServer(service)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		return terr.Wrap(err)
	}

	llog.Msg("Start suggester grpc server").Level(llog.INFO).Data("port", grpcPort).Log(ctx)

	grpcServer := grpc.NewServer(utils.Middlewares()...)
	suggester_grpc.RegisterSuggesterGrpcServer(grpcServer, suggesterGrpcserver)

	err = grpcServer.Serve(listener)
	if err != nil {
		return terr.Wrap(err)
	}

	return nil
}
