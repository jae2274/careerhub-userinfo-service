package restapi

// import (
// 	"context"
// 	"testing"

// 	"github.com/jae2274/careerhub-userinfo-service/careerhub/userinfo_service/restapi/restapi_grpc"
// 	"github.com/jae2274/careerhub-userinfo-service/test/tinit"
// 	"github.com/stretchr/testify/require"
// )

// func TestRestapiGrpc(t *testing.T) {
// 	cancelFunc := tinit.RunTestApp(t)
// 	defer cancelFunc()

// 	t.Run("return empty", func(t *testing.T) {
// 		ctx := context.Background()
// 		client := tinit.InitRestapiClient(t)
// 		res, err := client.FindMatchJob()(ctx, &restapi_grpc.GetScrapJobsRequest{
// 			UserId: "testUserId",
// 		})
// 		require.NoError(t, err)
// 		require.Empty(t, res.ScrapJobs)

// 	})
// }
