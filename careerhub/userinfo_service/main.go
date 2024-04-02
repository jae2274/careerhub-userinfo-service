package main

import (
	"context"
	"os"

	"github.com/jae2274/careerhub-userinfo-service/careerhub/userinfo_service/common/domain/condition"
	"github.com/jae2274/careerhub-userinfo-service/careerhub/userinfo_service/common/mongocfg"
	"github.com/jae2274/careerhub-userinfo-service/careerhub/userinfo_service/common/vars"
	"github.com/jae2274/careerhub-userinfo-service/careerhub/userinfo_service/restapi"
	"github.com/jae2274/goutils/llog"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	app     = "userinfo-service"
	service = "careerhub"

	ctxKeyTraceID = "trace_id"
)

func main() {
	ctx := context.Background()

	err := initLogger(ctx)
	checkErr(ctx, err)

	envVars, err := vars.Variables()
	checkErr(ctx, err)

	db, err := mongocfg.NewDatabase(envVars.MongoUri, envVars.DbName, envVars.DBUser)
	checkErr(ctx, err)

	collections, err := initCollections(db)
	checkErr(ctx, err)

	runErr := make(chan error)

	go func() {
		err := restapi.Run(ctx, envVars.RestApiGrpcPort, collections)
		runErr <- err
	}()

	err = <-runErr
	checkErr(ctx, err)
}

func initLogger(ctx context.Context) error {
	llog.SetMetadata("service", service)
	llog.SetMetadata("app", app)
	llog.SetDefaultContextData(ctxKeyTraceID)

	hostname, err := os.Hostname()
	if err != nil {
		return err
	}

	llog.SetMetadata("hostname", hostname)

	return nil
}

func initCollections(db *mongo.Database) (map[string]*mongo.Collection, error) {
	collections, err := mongocfg.InitCollections(db, &condition.DesiredCondition{})
	if err != nil {
		return nil, err
	}

	return collections, nil
}

func checkErr(ctx context.Context, err error) {
	if err != nil {
		llog.LogErr(ctx, err)
		os.Exit(1)
	}
}
