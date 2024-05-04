package scrapjob

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	UserIdField    = "userId"
	SiteField      = "site"
	PostingIdField = "postingId"
)

type ScrapJob struct {
	UserId    string `bson:"userId"`
	Site      string `bson:"site"`
	PostingId string `bson:"postingId"`
}

func (*ScrapJob) Collection() string {
	return "scrapJob"
}

func (*ScrapJob) IndexModels() map[string]*mongo.IndexModel {
	indexName := fmt.Sprintf("%s_1_%s_1_%s_1", UserIdField, SiteField, PostingIdField) // userId_1
	return map[string]*mongo.IndexModel{
		indexName: {
			Keys: bson.D{
				{Key: UserIdField, Value: 1},
				{Key: SiteField, Value: 1},
				{Key: PostingIdField, Value: 1},
			},
			Options: options.Index().SetUnique(true),
		},
	}
}
