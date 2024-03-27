package history

import (
	"time"

	"github.com/jae2274/goutils/enum"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type WorkStateValues struct{}

type WorkState = enum.Enum[WorkStateValues]

const (
	WORKING  = WorkState("WORKING")
	FINISHED = WorkState("FINISHED")
	FAILED   = WorkState("FAILED")
)

func (WorkStateValues) Values() []string {
	return []string{string(WORKING), string(FINISHED), string(FAILED)}
}

type SuggesterHistory struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	StartTime  time.Time          `bson:"startTime"`
	EndTime    time.Time          `bson:"endTime"`
	State      WorkState          `bson:"state"`
	InsertedAt time.Time          `bson:"insertedAt"`
}

func (*SuggesterHistory) Collection() string {
	return "suggesterHistory"
}

func (*SuggesterHistory) IndexModels() map[string]*mongo.IndexModel {
	return map[string]*mongo.IndexModel{}
}
