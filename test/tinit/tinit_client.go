package tinit

import (
	"testing"

	"github.com/jae2274/careerhub-userinfo-service/careerhub/userinfo_service/restapi/restapi_grpc"
	"github.com/jae2274/careerhub-userinfo-service/careerhub/userinfo_service/suggester/suggester_grpc"
)

func InitSuggesterClient(t *testing.T) suggester_grpc.UserinfoClient {
	envVars := InitEnvVars(t)
	conn := InitGrpcClient(t, envVars.SuggesterGrpcPort)

	return suggester_grpc.NewUserinfoClient(conn)
}

func InitRestapiClient(t *testing.T) restapi_grpc.MatchJobGrpcClient {
	envVars := InitEnvVars(t)
	conn := InitGrpcClient(t, envVars.MatchJobGrpcPort)

	return restapi_grpc.NewMatchJobGrpcClient(conn)
}