package repo

import (
	"context"
	"time"
)

type HistoryRepo interface {
	GetLastSuccessDate(context.Context) (time.Time, error)
	InsertHistory(context.Context, string) error
	UpdateHistory(context.Context, string) error
}
