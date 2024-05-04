package repo

import (
	"context"

	"github.com/jae2274/careerhub-userinfo-service/careerhub/userinfo_service/common/domain/scrapjob"
	"github.com/jae2274/goutils/terr"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ScrapJobRepo interface {
	GetScrapJobs(ctx context.Context, userId string, tag *string) ([]*scrapjob.ScrapJob, error)
	AddScrapJob(ctx context.Context, scrapJob *scrapjob.ScrapJob) error
	RemoveScrapJob(ctx context.Context, userId, site, postingId string) error
	AddTag(ctx context.Context, userId, site, postingId, tag string) error
	RemoveTag(ctx context.Context, userId, site, postingId, tag string) error
	GetScrapTags(ctx context.Context, userId string) ([]string, error)
}

type ScrapJobRepoImpl struct {
	col *mongo.Collection
}

func NewScrapJobRepo(db *mongo.Database) ScrapJobRepo {
	return &ScrapJobRepoImpl{
		col: db.Collection((&scrapjob.ScrapJob{}).Collection()),
	}
}

func (r *ScrapJobRepoImpl) GetScrapJobs(ctx context.Context, userId string, tag *string) ([]*scrapjob.ScrapJob, error) {
	filter := bson.M{scrapjob.UserIdField: userId}
	if tag != nil {
		filter[scrapjob.TagsField] = *tag
	}
	cur, err := r.col.Find(ctx, filter)
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

func (r *ScrapJobRepoImpl) AddTag(ctx context.Context, userId, site, postingId, tag string) error {
	_, err := r.col.UpdateOne(ctx, bson.M{
		scrapjob.UserIdField:    userId,
		scrapjob.SiteField:      site,
		scrapjob.PostingIdField: postingId,
	}, bson.M{
		"$addToSet": bson.M{scrapjob.TagsField: tag},
	})

	if err != nil {
		return terr.Wrap(err)
	}

	return nil
}

func (r *ScrapJobRepoImpl) RemoveTag(ctx context.Context, userId, site, postingId, tag string) error {
	_, err := r.col.UpdateOne(ctx, bson.M{
		scrapjob.UserIdField:    userId,
		scrapjob.SiteField:      site,
		scrapjob.PostingIdField: postingId,
	}, bson.M{
		"$pull": bson.M{scrapjob.TagsField: tag},
	})

	if err != nil {
		return terr.Wrap(err)
	}

	return nil
}

func (r *ScrapJobRepoImpl) GetScrapTags(ctx context.Context, userId string) ([]string, error) {
	results, err := r.col.Distinct(ctx, scrapjob.TagsField, bson.M{scrapjob.UserIdField: userId})

	if err != nil {
		return nil, terr.Wrap(err)
	}

	tags := make([]string, len(results))
	for i, result := range results {
		tag, ok := result.(string)
		if !ok {
			return nil, terr.New("invalid tags type")
		}

		tags[i] = tag
	}

	return tags, nil
}
