package repo

import (
	"context"
	"time"

	"github.com/jae2274/careerhub-userinfo-service/careerhub/userinfo_service/common/domain/history"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type HistoryRepo interface {
	FindWorkingBatchIds(context.Context) ([]string, error)
	UpdateFailed(context.Context, []string) error
	FindLastSuccessedDate(context.Context) (*time.Time, error)
	InsertHistory(context.Context, string, time.Time) error
	UpdateSuccessed(context.Context, string, time.Time) error
}

type HistoryRepoImpl struct {
	col *mongo.Collection
}

func NewHistoryRepo(db *mongo.Database) HistoryRepo {
	return &HistoryRepoImpl{
		col: db.Collection((&history.History{}).Collection()),
	}
}

func (r *HistoryRepoImpl) FindWorkingBatchIds(ctx context.Context) ([]string, error) {
	filter := bson.M{history.BatchStateField: history.WORKING}
	option := options.Find().SetProjection(bson.M{history.BatchIdField: 1})
	cursor, err := r.col.Find(ctx, filter, option)
	if err != nil {
		if mongo.ErrNilDocument == err {
			return []string{}, nil
		}
		return nil, err
	}
	defer cursor.Close(ctx)

	var batchIds []history.History
	err = cursor.All(ctx, &batchIds)
	if err != nil {
		return nil, err
	}

	ids := make([]string, len(batchIds))
	for i, batchId := range batchIds {
		ids[i] = batchId.BatchId
	}

	return ids, nil
}

func (r *HistoryRepoImpl) UpdateFailed(ctx context.Context, batchIds []string) error {
	filter := bson.M{history.BatchIdField: bson.M{"$in": batchIds}}
	update := bson.M{"$set": bson.M{history.BatchStateField: history.FAILED, history.EndTimeField: time.Now()}}
	_, err := r.col.UpdateMany(ctx, filter, update)
	return err
}

func (r *HistoryRepoImpl) FindLastSuccessedDate(ctx context.Context) (*time.Time, error) {
	filter := bson.M{history.BatchStateField: history.SUCCESSED}
	option := options.FindOne().SetSort(bson.M{history.EndTimeField: -1})
	var history history.History
	err := r.col.FindOne(ctx, filter, option).Decode(&history)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return history.EndTime, nil
}

func (r *HistoryRepoImpl) InsertHistory(ctx context.Context, batchId string, startTime time.Time) error {
	history := &history.History{
		BatchId:    batchId,
		BatchState: history.WORKING,
		StartTime:  startTime,
		InsertedAt: time.Now(),
	}
	_, err := r.col.InsertOne(ctx, history)
	return err
}

func (r *HistoryRepoImpl) UpdateSuccessed(ctx context.Context, batchId string, endTime time.Time) error {
	filter := bson.M{history.BatchIdField: batchId}
	update := bson.M{"$set": bson.M{history.BatchStateField: history.SUCCESSED, history.EndTimeField: endTime}}
	_, err := r.col.UpdateOne(ctx, filter, update)
	return err
}
