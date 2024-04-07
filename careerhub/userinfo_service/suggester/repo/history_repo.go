package repo

import (
	"context"
	"time"
)

type HistoryRepo interface {
	FindWorkingBatchIds(context.Context) ([]string, error)
	UpdateFailed(context.Context, []string) error
	FindLastSuccessedDate(context.Context) (*time.Time, error)
	InsertHistory(context.Context, string) error
	UpdateSuccessed(context.Context, string) error
}
