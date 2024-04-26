package repo

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jae2274/careerhub-userinfo-service/careerhub/userinfo_service/common/domain/matchjob"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MatchJobRepo interface {
	InitMatchJob(ctx context.Context, userId string) (bool, error)
	FindByUserId(ctx context.Context, userId string) (*matchjob.MatchJob, error)
	InsertCondition(ctx context.Context, userId string, limitCount uint, newCondition *matchjob.Condition) (bool, error)
	UpdateCondition(ctx context.Context, userId string, updateCondition *matchjob.Condition) (bool, error)
	DeleteCondition(ctx context.Context, userId string, conditionId string) (bool, error)
	UpdateAgreeToMail(ctx context.Context, userId string, agreeToMail bool) (bool, error)
}

type MatchJobRepoImpl struct {
	col *mongo.Collection
}

func NewMatchJobRepo(db *mongo.Database) MatchJobRepo {
	return &MatchJobRepoImpl{
		col: db.Collection((&matchjob.MatchJob{}).Collection()),
	}
}

func (r *MatchJobRepoImpl) InitMatchJob(ctx context.Context, userId string) (bool, error) {
	condition := matchjob.MatchJob{
		UserId:     userId,
		Conditions: []*matchjob.Condition{},
		InsertedAt: time.Now(),
	}

	_, err := r.col.InsertOne(ctx, condition)

	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func (r *MatchJobRepoImpl) FindByUserId(ctx context.Context, userId string) (*matchjob.MatchJob, error) {
	filter := bson.M{matchjob.UserIdField: userId}

	matchJob := &matchjob.MatchJob{}
	err := r.col.FindOne(ctx, filter).Decode(matchJob)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}

		return nil, err
	}

	return matchJob, nil
}

var ErrNonZero = fmt.Errorf("limitCount must be greater than 0")

func (r *MatchJobRepoImpl) InsertCondition(ctx context.Context, userId string, limitCount uint, newCondition *matchjob.Condition) (bool, error) {
	if limitCount == 0 {
		return false, ErrNonZero
	}
	newCondition.ConditionId = uuid.NewString()

	filter := bson.M{
		matchjob.UserIdField: userId, //해당 조건은 InitConditions 함수에서 생성되므로 userId가 존재한다는 것을 보장함
		fmt.Sprintf("%s.%d", matchjob.ConditionsField, limitCount-1): bson.M{"$exists": false}, //갯수 제한
	}
	update := bson.M{"$push": bson.M{matchjob.ConditionsField: newCondition}}

	result, err := r.col.UpdateOne(ctx, filter, update)

	if err != nil {
		return false, err
	}

	if result.MatchedCount == 0 {
		return false, nil
	}

	return true, nil
}

func (r *MatchJobRepoImpl) UpdateCondition(ctx context.Context, userId string, updateCondition *matchjob.Condition) (bool, error) {
	filter := bson.M{
		matchjob.UserIdField:                 userId,
		matchjob.Conditions_ConditionIdField: updateCondition.ConditionId,
	}
	update := bson.M{"$set": bson.M{fmt.Sprintf("%s.$", matchjob.ConditionsField): updateCondition}}

	result, err := r.col.UpdateOne(ctx, filter, update)

	if err != nil {
		return false, err
	}

	if result.MatchedCount == 0 {
		return false, nil
	}

	return true, nil
}

func (r *MatchJobRepoImpl) DeleteCondition(ctx context.Context, userId string, conditionId string) (bool, error) {
	filter := bson.M{
		matchjob.UserIdField:                 userId,
		matchjob.Conditions_ConditionIdField: conditionId,
	}
	update := bson.M{"$pull": bson.M{matchjob.ConditionsField: bson.M{matchjob.ConditionIdField: conditionId}}}

	result, err := r.col.UpdateOne(ctx, filter, update)

	if err != nil {
		return false, err
	}

	if result.MatchedCount == 0 {
		return false, nil
	}

	return true, nil
}

func (r *MatchJobRepoImpl) UpdateAgreeToMail(ctx context.Context, userId string, agreeToMail bool) (bool, error) {
	filter := bson.M{matchjob.UserIdField: userId}
	update := bson.M{"$set": bson.M{matchjob.AgreeToMailField: agreeToMail}}

	result, err := r.col.UpdateOne(ctx, filter, update)
	if err != nil {
		return false, err
	}

	if result.MatchedCount == 0 {
		return false, nil
	}

	return true, nil
}
