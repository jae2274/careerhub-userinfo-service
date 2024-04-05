package restapi

import (
	"context"
	"fmt"
	"net"

	"github.com/jae2274/careerhub-userinfo-service/careerhub/userinfo_service/common/domain/condition"
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

func Run(ctx context.Context, grpcPort int, collections map[string]*mongo.Collection) error {
	conditionRepo := repo.NewConditionRepo(collections[(&condition.DesiredCondition{}).Collection()])
	conditionService := service.NewConditionService(conditionRepo)
	restApiService := server.NewRestApiGrpcServer(conditionService)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		return terr.Wrap(err)
	}

	llog.Msg("Start restapi grpc server").Level(llog.INFO).Data("port", grpcPort).Log(ctx)

	grpcServer := grpc.NewServer(utils.Middlewares()...)
	restapi_grpc.RegisterRestApiGrpcServer(grpcServer, restApiService)

	err = grpcServer.Serve(listener)
	if err != nil {
		return terr.Wrap(err)
	}

	return nil
}
