package repo

import (
	"context"

	condition "github.com/jae2274/careerhub-userinfo-service/careerhub/userinfo_service/common/domain/matchjob"
	"github.com/jae2274/goutils/terr"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MatchJobRepo interface {
	GetMatchJobs(context.Context) ([]*condition.MatchJob, error)
}

type MatchJobRepoImpl struct {
	col *mongo.Collection
}

func NewMatchJobRepo(db *mongo.Database) MatchJobRepo {
	return &MatchJobRepoImpl{
		col: db.Collection((&condition.MatchJob{}).Collection()),
	}
}

func (r *MatchJobRepoImpl) GetMatchJobs(ctx context.Context) ([]*condition.MatchJob, error) {
	cursor, err := r.col.Find(ctx, bson.M{condition.AgreeToMailField: true})
	if err != nil {
		if mongo.ErrNilDocument == err {
			return []*condition.MatchJob{}, nil
		}
		return nil, terr.Wrap(err)
	}
	defer cursor.Close(ctx)

	var matchJobs []*condition.MatchJob
	err = cursor.All(ctx, &matchJobs)
	if err != nil {
		return nil, terr.Wrap(err)
	}

	return matchJobs, nil
}
