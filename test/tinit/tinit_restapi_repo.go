package tinit

import (
	"testing"

	"github.com/jae2274/careerhub-userinfo-service/careerhub/userinfo_service/common/domain/condition"
	"github.com/jae2274/careerhub-userinfo-service/careerhub/userinfo_service/restapi/repo"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitRestapiConditionRepo(t *testing.T, db *mongo.Database) repo.ConditionRepo {
	col := db.Collection((&condition.DesiredCondition{}).Collection())
	repo := repo.NewConditionRepo(col)

	return repo
}
