package condition

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	UserIdField                 = "userId"
	Conditions_ConditionIdField = "conditions.conditionId"
	ConditionIdField            = "conditionId"
	ConditionsField             = "conditions"
)

type DesiredCondition struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	UserId     string             `bson:"userId"`
	Conditions []Condition        `bson:"conditions"`
	InsertedAt time.Time          `bson:"insertedAt"`
}

type Condition struct {
	ConditionId   string `bson:"conditionId"`
	ConditionName string `bson:"conditionName"`
	Query         Query  `bson:"query"`
}

type Query struct {
	Categories []*CategoryQuery `bson:"categories"`
	SkillNames [][]string       `bson:"skillNames"`
	MinCareer  *int32           `bson:"minCareer"`
	MaxCareer  *int32           `bson:"maxCareer"`
}

type CategoryQuery struct {
	Site         string `bson:"site"`
	CategoryName string `bson:"categoryName"`
}

func (*DesiredCondition) Collection() string {
	return "desiredCondition"
}

func (*DesiredCondition) IndexModels() map[string]*mongo.IndexModel {
	useridIndex := fmt.Sprintf("%s_1", UserIdField)                                // userId_1
	conditionsConditionIdIndex := fmt.Sprintf("%s_1", Conditions_ConditionIdField) // conditions.conditionId_1
	return map[string]*mongo.IndexModel{
		useridIndex: {
			Keys:    bson.D{{Key: UserIdField, Value: 1}},
			Options: options.Index().SetUnique(true),
		},
		conditionsConditionIdIndex: {
			Keys:    bson.D{{Key: Conditions_ConditionIdField, Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	}
}
