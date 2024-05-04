package app

import (
	"context"
	"os"

	"github.com/jae2274/careerhub-userinfo-service/careerhub/userinfo_service/common/domain/matchjob"
	"github.com/jae2274/careerhub-userinfo-service/careerhub/userinfo_service/common/mongocfg"
	"github.com/jae2274/careerhub-userinfo-service/careerhub/userinfo_service/common/vars"
	"github.com/jae2274/careerhub-userinfo-service/careerhub/userinfo_service/restapi"
	"github.com/jae2274/careerhub-userinfo-service/careerhub/userinfo_service/suggester"
	"github.com/jae2274/goutils/llog"
	"github.com/jae2274/goutils/mw"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	app     = "userinfo-service"
	service = "careerhub"

	ctxKeyTraceID = string(mw.CtxKeyTraceID)
)

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

func Run(ctx context.Context) {

	err := initLogger(ctx)
	checkErr(ctx, err)
	llog.Info(ctx, "Start Application")

	envVars, err := vars.Variables()
	checkErr(ctx, err)

	db, err := mongocfg.NewDatabase(envVars.MongoUri, envVars.DbName, envVars.DBUser)
	checkErr(ctx, err)

	err = initCollections(db)
	checkErr(ctx, err)

	runErr := make(chan error)

	go func() {
		err := restapi.Run(ctx, envVars.MatchJobGrpcPort, db)
		runErr <- err
	}()

	go func() {
		err := suggester.Run(ctx, envVars.SuggesterGrpcPort, db)
		runErr <- err
	}()

	select {
	case <-ctx.Done():
		llog.Info(ctx, "Finish Application")
	case err := <-runErr:
		checkErr(ctx, err)
	}
}

func initCollections(db *mongo.Database) error {
	_, err := mongocfg.InitCollections(db, &matchjob.MatchJob{})
	if err != nil {
		return err
	}

	return nil
}

func checkErr(ctx context.Context, err error) {
	if err != nil {
		llog.LogErr(ctx, err)
		os.Exit(1)
	}
}
