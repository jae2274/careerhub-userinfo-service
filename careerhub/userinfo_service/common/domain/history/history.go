package history

import (
	"fmt"
	"time"

	"github.com/jae2274/goutils/enum"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	BatchIdField = "batchId"
)

type BatchStateValues struct{}

type BatchState = enum.Enum[BatchStateValues]

const (
	WORKING   = BatchState("WORKING")
	SUCCESSED = BatchState("SUCCESSED")
	FAILED    = BatchState("FAILED")
)

func (BatchStateValues) Values() []string {
	return []string{string(WORKING), string(SUCCESSED), string(FAILED)}
}

type History struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`

	BatchId    string     `bson:"batchId"`
	BatchState BatchState `bson:"batchState"`
	StartTime  time.Time  `bson:"startTime"`
	EndTime    *time.Time `bson:"endTime"`

	InsertedAt time.Time `bson:"insertedAt"`
}

func (*History) Collection() string {
	return "history"
}

func (*History) IndexModels() map[string]*mongo.IndexModel {
	batchIdIndex := fmt.Sprintf("%s_1", BatchIdField)
	return map[string]*mongo.IndexModel{
		batchIdIndex: {
			Keys:    bson.D{{Key: BatchIdField, Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	}
}
