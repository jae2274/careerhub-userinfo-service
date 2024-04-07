package repo

import (
	"context"

	"github.com/jae2274/careerhub-userinfo-service/careerhub/userinfo_service/common/domain/condition"
)

type ConditionRepo interface {
	GetDesiredConditions(context.Context) ([]condition.DesiredCondition, error)
}
