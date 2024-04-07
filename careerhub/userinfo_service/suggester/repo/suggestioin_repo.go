package repo

import (
	"context"

	"github.com/jae2274/careerhub-userinfo-service/careerhub/userinfo_service/common/domain/suggestion"
)

type SuggestionRepo interface {
	DeleteSuggestions(context.Context, []string) error
	InsertSuggestion(context.Context, *suggestion.Suggestion) error
}
