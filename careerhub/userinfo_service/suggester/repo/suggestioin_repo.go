package repo

import (
	"context"

	"github.com/jae2274/careerhub-userinfo-service/careerhub/userinfo_service/common/domain/suggestion"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type SuggestionRepo interface {
	DeleteSuggestions(context.Context, []string) error
	InsertSuggestion(context.Context, *suggestion.Suggestion) error
}

type SuggestionRepoImpl struct {
	col *mongo.Collection
}

func NewSuggestionRepo(db *mongo.Database) SuggestionRepo {
	return &SuggestionRepoImpl{
		col: db.Collection((&suggestion.Suggestion{}).Collection()),
	}
}

func (r *SuggestionRepoImpl) DeleteSuggestions(ctx context.Context, batchIds []string) error {
	filter := bson.M{suggestion.BatchIdField: bson.M{"$in": batchIds}}
	_, err := r.col.DeleteMany(ctx, filter)
	return err
}

func (r *SuggestionRepoImpl) InsertSuggestion(ctx context.Context, suggestion *suggestion.Suggestion) error {
	_, err := r.col.InsertOne(ctx, suggestion)
	return err
}
