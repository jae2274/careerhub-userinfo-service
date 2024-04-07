package suggestion

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	BatchIdField     = "batchId"
	UserIdField      = "userId"
	ConditionIdField = "conditionId"
)

type Suggestion struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	BatchId       string             `bson:"batchId"`
	UserId        string             `bson:"userId"`
	ConditionId   string             `bson:"conditionId"`
	ConditionName string             `bson:"conditionName"`
	Postings      []*Posting         `bson:"postings"`
	InsertedAt    time.Time          `bson:"insertedAt"`
}

type PostingId struct {
	Site      string `bson:"site"`
	PostingId string `bson:"postingId"`
}

type Posting struct {
	PostingId   *PostingId   `bson:"postingId"`
	Title       string       `bson:"title"`
	CompanyId   string       `bson:"companyId"`
	CompanyName string       `bson:"companyName"`
	Info        *PostingInfo `bson:"info"`
}

type PostingInfo struct {
	Categories []string `bson:"categories"`
	SkillNames []string `bson:"skillNames"`
	MinCareer  *int32   `bson:"minCareer"`
	MaxCareer  *int32   `bson:"maxCareer"`
}

func (*Suggestion) Collection() string {
	return "suggestion"
}

func (*Suggestion) IndexModels() map[string]*mongo.IndexModel {
	batchIduserIdIndex := BatchIdField + "_1_" + UserIdField + "_1"
	return map[string]*mongo.IndexModel{
		batchIduserIdIndex: {
			Keys:    bson.D{{Key: BatchIdField, Value: 1}, {Key: UserIdField, Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	}
}
