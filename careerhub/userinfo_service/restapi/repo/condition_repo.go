package repo

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jae2274/careerhub-userinfo-service/careerhub/userinfo_service/common/domain/condition"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ConditionRepo interface {
	InitDesiredCondition(ctx context.Context, userId string) (bool, error)
	FindByUserId(ctx context.Context, userId string) (*condition.DesiredCondition, error)
	InsertCondition(ctx context.Context, userId string, limitCount uint, newCondition *condition.Condition) (bool, error)
	UpdateCondition(ctx context.Context, userId string, updateCondition *condition.Condition) (bool, error)
	DeleteCondition(ctx context.Context, userId string, conditionId string) (bool, error)
	UpdateAgreeToMail(ctx context.Context, userId string, agreeToMail bool) (bool, error)
}

type ConditionRepoImpl struct {
	col *mongo.Collection
}

func NewConditionRepo(db *mongo.Database) ConditionRepo {
	return &ConditionRepoImpl{
		col: db.Collection((&condition.DesiredCondition{}).Collection()),
	}
}

func (r *ConditionRepoImpl) InitDesiredCondition(ctx context.Context, userId string) (bool, error) {
	condition := condition.DesiredCondition{
		UserId:     userId,
		Conditions: []*condition.Condition{},
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

func (r *ConditionRepoImpl) FindByUserId(ctx context.Context, userId string) (*condition.DesiredCondition, error) {
	filter := bson.M{condition.UserIdField: userId}

	desiredCondition := &condition.DesiredCondition{}
	err := r.col.FindOne(ctx, filter).Decode(desiredCondition)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}

		return nil, err
	}

	return desiredCondition, nil
}

var ErrNonZero = fmt.Errorf("limitCount must be greater than 0")

func (r *ConditionRepoImpl) InsertCondition(ctx context.Context, userId string, limitCount uint, newCondition *condition.Condition) (bool, error) {
	if limitCount == 0 {
		return false, ErrNonZero
	}
	newCondition.ConditionId = uuid.NewString()

	filter := bson.M{
		condition.UserIdField: userId, //해당 조건은 InitConditions 함수에서 생성되므로 userId가 존재한다는 것을 보장함
		fmt.Sprintf("%s.%d", condition.ConditionsField, limitCount-1): bson.M{"$exists": false},                //갯수 제한
		condition.Conditions_ConditionIdField:                         bson.M{"$ne": newCondition.ConditionId}, //중복 방지
	}
	update := bson.M{"$push": bson.M{condition.ConditionsField: newCondition}}

	result, err := r.col.UpdateOne(ctx, filter, update)

	if err != nil {
		return false, err
	}

	if result.MatchedCount == 0 {
		return false, nil
	}

	return true, nil
}

func (r *ConditionRepoImpl) UpdateCondition(ctx context.Context, userId string, updateCondition *condition.Condition) (bool, error) {
	filter := bson.M{
		condition.UserIdField:                 userId,
		condition.Conditions_ConditionIdField: updateCondition.ConditionId,
	}
	update := bson.M{"$set": bson.M{fmt.Sprintf("%s.$", condition.ConditionsField): updateCondition}}

	result, err := r.col.UpdateOne(ctx, filter, update)

	if err != nil {
		return false, err
	}

	if result.MatchedCount == 0 {
		return false, nil
	}

	return true, nil
}

func (r *ConditionRepoImpl) DeleteCondition(ctx context.Context, userId string, conditionId string) (bool, error) {
	filter := bson.M{
		condition.UserIdField:                 userId,
		condition.Conditions_ConditionIdField: conditionId,
	}
	update := bson.M{"$pull": bson.M{condition.ConditionsField: bson.M{condition.ConditionIdField: conditionId}}}

	result, err := r.col.UpdateOne(ctx, filter, update)

	if err != nil {
		return false, err
	}

	if result.MatchedCount == 0 {
		return false, nil
	}

	return true, nil
}

func (r *ConditionRepoImpl) UpdateAgreeToMail(ctx context.Context, userId string, agreeToMail bool) (bool, error) {
	filter := bson.M{condition.UserIdField: userId}
	update := bson.M{"$set": bson.M{condition.AgreeToMailField: agreeToMail}}

	result, err := r.col.UpdateOne(ctx, filter, update)
	if err != nil {
		return false, err
	}

	if result.MatchedCount == 0 {
		return false, nil
	}

	return true, nil
}
