package repo

import (
	"context"

	"github.com/jae2274/careerhub-userinfo-service/careerhub/userinfo_service/common/domain/scrapjob"
	"github.com/jae2274/goutils/terr"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ScrapJobRepo interface {
	GetScrapJobs(ctx context.Context, userId string) ([]*scrapjob.ScrapJob, error)
	AddScrapJob(ctx context.Context, scrapJob *scrapjob.ScrapJob) error
	RemoveScrapJob(ctx context.Context, userId, site, postingId string) error
}

type ScrapJobRepoImpl struct {
	col *mongo.Collection
}

func NewScrapJobRepo(db *mongo.Database) ScrapJobRepo {
	return &ScrapJobRepoImpl{
		col: db.Collection((&scrapjob.ScrapJob{}).Collection()),
	}
}

func (r *ScrapJobRepoImpl) GetScrapJobs(ctx context.Context, userId string) ([]*scrapjob.ScrapJob, error) {
	cur, err := r.col.Find(ctx, bson.M{scrapjob.UserIdField: userId})
	if err != nil {
		return nil, terr.Wrap(err)
	}

	var scrapJobs []*scrapjob.ScrapJob
	err = cur.All(ctx, &scrapJobs)
	if err != nil {
		return nil, terr.Wrap(err)
	}

	return scrapJobs, nil
}

func (r *ScrapJobRepoImpl) AddScrapJob(ctx context.Context, scrapJob *scrapjob.ScrapJob) error {
	_, err := r.col.InsertOne(ctx, scrapJob)
	if err != nil {
		return terr.Wrap(err)
	}
	return nil
}

func (r *ScrapJobRepoImpl) RemoveScrapJob(ctx context.Context, userId, site, postingId string) error {
	_, err := r.col.DeleteOne(ctx, bson.M{
		scrapjob.UserIdField:    userId,
		scrapjob.SiteField:      site,
		scrapjob.PostingIdField: postingId,
	})
	if err != nil {
		return terr.Wrap(err)
	}
	return nil
}
