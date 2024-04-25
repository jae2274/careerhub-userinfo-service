package repo

import (
	"context"

	"github.com/jae2274/careerhub-userinfo-service/careerhub/userinfo_service/common/domain/condition"
	"github.com/jae2274/goutils/terr"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ConditionRepo interface {
	GetDesiredConditions(context.Context) ([]*condition.DesiredCondition, error)
}

type ConditionRepoImpl struct {
	col *mongo.Collection
}

func NewConditionRepo(db *mongo.Database) ConditionRepo {
	return &ConditionRepoImpl{
		col: db.Collection((&condition.DesiredCondition{}).Collection()),
	}
}

func (r *ConditionRepoImpl) GetDesiredConditions(ctx context.Context) ([]*condition.DesiredCondition, error) {
	cursor, err := r.col.Find(ctx, bson.M{condition.AgreeToMailField: true})
	if err != nil {
		if mongo.ErrNilDocument == err {
			return []*condition.DesiredCondition{}, nil
		}
		return nil, terr.Wrap(err)
	}
	defer cursor.Close(ctx)

	var desiredConditions []*condition.DesiredCondition
	err = cursor.All(ctx, &desiredConditions)
	if err != nil {
		return nil, terr.Wrap(err)
	}

	return desiredConditions, nil
}
