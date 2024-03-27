package repo

import (
	"context"
	"fmt"
	"time"

	"github.com/jae2274/careerhub-userinfo-service/careerhub/userinfo_service/common/domain/condition"
	"github.com/jae2274/goutils/terr"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ConditionRepo interface {
	InitConditions(ctx context.Context, userId string) error
	FindByUserId(ctx context.Context, userId string) (*condition.DesiredCondition, error)
	FindByUserIdAndUUID(ctx context.Context, userId string, conditionId string) (*condition.Condition, error)
	InsertCondition(ctx context.Context, userId string, limitCount int, newCondition *condition.Condition) (bool, error)
	UpdateCondition(ctx context.Context, userId string, updateCondition *condition.Condition) (bool, error)
	DeleteCondition(ctx context.Context, userId string, conditionId string) (bool, error)
}

type ConditionRepoImpl struct {
	col *mongo.Collection
}

func NewConditionRepo(col *mongo.Collection) ConditionRepo {
	return &ConditionRepoImpl{
		col: col,
	}
}

func (r *ConditionRepoImpl) InitConditions(ctx context.Context, userId string) error {
	condition := condition.DesiredCondition{
		UserId:     userId,
		InsertedAt: time.Now(),
	}

	_, err := r.col.InsertOne(ctx, condition)

	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return nil
		}

		return err
	}

	return nil
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

func (r *ConditionRepoImpl) FindByUserIdAndUUID(ctx context.Context, userId string, conditionId string) (*condition.Condition, error) {
	filter := bson.M{condition.UserIdField: userId, condition.Conditions_ConditionIdField: conditionId}

	condition := &condition.DesiredCondition{}
	err := r.col.FindOne(ctx, filter).Decode(condition)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}

		return nil, err
	}

	for _, c := range condition.Conditions {
		if c.ConditionId == conditionId {
			return &c, nil
		}
	}

	//정상 동작 시 이곳에 도달하지 않음
	return nil, terr.New(fmt.Sprintf("condition not found: userId=%s, conditionId=%s", userId, conditionId))
}

func (r *ConditionRepoImpl) InsertCondition(ctx context.Context, userId string, limitCount int, newCondition *condition.Condition) (bool, error) {
	filter := bson.M{
		condition.UserIdField: userId, //해당 조건은 InitConditions 함수에서 생성되므로 userId가 존재한다는 것을 보장함
		fmt.Sprintf("%s.%d", condition.ConditionsField, limitCount): bson.M{"$exists": false}, //갯수 제한
	}
	update := bson.M{"$push": bson.M{condition.Conditions_ConditionIdField: newCondition}}

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
	update := bson.M{"$pull": bson.M{condition.Conditions_ConditionIdField: bson.M{"conditionId": conditionId}}}

	result, err := r.col.UpdateOne(ctx, filter, update)

	if err != nil {
		return false, err
	}

	if result.MatchedCount == 0 {
		return false, nil
	}

	return true, nil
}
