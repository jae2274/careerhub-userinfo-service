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
	WORKING  = BatchState("WORKING")
	FINISHED = BatchState("FINISHED")
	FAILED   = BatchState("FAILED")
)

func (BatchStateValues) Values() []string {
	return []string{string(WORKING), string(FINISHED), string(FAILED)}
}

type MailerStateValues struct{}

type MailerState = enum.Enum[MailerStateValues]

const (
	NOT_SENT = MailerState("NOT_SENT")
	ALL_SENT = MailerState("ALL_SENT")
)

func (MailerStateValues) Values() []string {
	return []string{string(NOT_SENT), string(ALL_SENT)}
}

type History struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`

	BatchId    string     `bson:"batchId"`
	BatchState BatchState `bson:"batchState"`
	StartTime  time.Time  `bson:"startTime"`
	EndTime    *time.Time `bson:"endTime"`

	MailerState MailerState `bson:"mailerState"`
	InsertedAt  time.Time   `bson:"insertedAt"`
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
