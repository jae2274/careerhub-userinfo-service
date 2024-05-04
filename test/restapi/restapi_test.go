package restapi

import (
	"context"
	"fmt"
	"testing"

	"github.com/jae2274/careerhub-userinfo-service/careerhub/userinfo_service/restapi/restapi_grpc"
	"github.com/jae2274/careerhub-userinfo-service/test/tinit"
	"github.com/stretchr/testify/require"
)

func TestRestapiGrpc(t *testing.T) {
	cancelFunc := tinit.RunTestApp(t)
	defer cancelFunc()

	t.Run("return empty", func(t *testing.T) {
		tinit.InitDB(t)
		ctx := context.Background()
		client := tinit.InitScrapJobGrpcClient(t)
		res, err := client.GetScrapJobs(ctx, &restapi_grpc.GetScrapJobsRequest{
			UserId: "testUserId",
		})

		require.NoError(t, err)
		require.Empty(t, res.ScrapJobs)

	})

	t.Run("return one scrapJob", func(t *testing.T) {
		tinit.InitDB(t)
		ctx := context.Background()
		client := tinit.InitScrapJobGrpcClient(t)

		userId := "testUserId"
		req := newAddScrapJobRequest(userId, 1)
		_, err := client.AddScrapJob(ctx, newAddScrapJobRequest(userId, 1))
		require.NoError(t, err)

		res, err := client.GetScrapJobs(ctx, &restapi_grpc.GetScrapJobsRequest{UserId: userId})

		require.NoError(t, err)
		require.Len(t, res.ScrapJobs, 1)
		require.Equal(t, req.Site, res.ScrapJobs[0].Site)
		require.Equal(t, req.PostingId, res.ScrapJobs[0].PostingId)
	})

	t.Run("return scrapJobs", func(t *testing.T) {
		tinit.InitDB(t)
		ctx := context.Background()
		client := tinit.InitScrapJobGrpcClient(t)

		userId := "testUserId"
		reqs := []*restapi_grpc.AddScrapJobRequest{
			newAddScrapJobRequest(userId, 1),
			newAddScrapJobRequest(userId, 2),
		}
		_, err := client.AddScrapJob(ctx, reqs[0])
		require.NoError(t, err)

		_, err = client.AddScrapJob(ctx, reqs[1])
		require.NoError(t, err)

		res, err := client.GetScrapJobs(ctx, &restapi_grpc.GetScrapJobsRequest{UserId: userId})

		require.NoError(t, err)
		require.Len(t, res.ScrapJobs, len(reqs))
		for i, req := range reqs {
			require.Equal(t, req.Site, res.ScrapJobs[i].Site)
			require.Equal(t, req.PostingId, res.ScrapJobs[i].PostingId)
		}
	})

	t.Run("return scrapJobs with different userId", func(t *testing.T) {
		tinit.InitDB(t)
		ctx := context.Background()
		client := tinit.InitScrapJobGrpcClient(t)

		userId1 := "testUserId1"
		userId2 := "testUserId2"
		reqs1 := []*restapi_grpc.AddScrapJobRequest{
			newAddScrapJobRequest(userId1, 1),
			newAddScrapJobRequest(userId1, 2),
		}
		reqs2 := []*restapi_grpc.AddScrapJobRequest{
			newAddScrapJobRequest(userId2, 1),
			newAddScrapJobRequest(userId2, 2),
		}

		for _, req := range append(reqs1, reqs2...) {
			_, err := client.AddScrapJob(ctx, req)
			require.NoError(t, err)
		}

		res1, err := client.GetScrapJobs(ctx, &restapi_grpc.GetScrapJobsRequest{UserId: userId1})

		require.NoError(t, err)
		require.Len(t, res1.ScrapJobs, len(reqs1))
		for i, req := range reqs1 {
			require.Equal(t, req.Site, res1.ScrapJobs[i].Site)
			require.Equal(t, req.PostingId, res1.ScrapJobs[i].PostingId)
		}

		res2, err := client.GetScrapJobs(ctx, &restapi_grpc.GetScrapJobsRequest{UserId: userId2})

		require.NoError(t, err)
		require.Len(t, res2.ScrapJobs, len(reqs2))
		for i, req := range reqs2 {
			require.Equal(t, req.Site, res2.ScrapJobs[i].Site)
			require.Equal(t, req.PostingId, res2.ScrapJobs[i].PostingId)
		}
	})

	t.Run("return scrapJobs without removed", func(t *testing.T) {
		tinit.InitDB(t)
		ctx := context.Background()
		client := tinit.InitScrapJobGrpcClient(t)

		userId := "testUserId"
		reqs := []*restapi_grpc.AddScrapJobRequest{
			newAddScrapJobRequest(userId, 1),
			newAddScrapJobRequest(userId, 2),
		}

		for _, req := range reqs {
			_, err := client.AddScrapJob(ctx, req)
			require.NoError(t, err)
		}

		_, err := client.RemoveScrapJob(ctx, &restapi_grpc.RemoveScrapJobRequest{
			UserId:    userId,
			Site:      reqs[0].Site,
			PostingId: reqs[0].PostingId,
		})
		require.NoError(t, err)

		res, err := client.GetScrapJobs(ctx, &restapi_grpc.GetScrapJobsRequest{UserId: userId})

		require.NoError(t, err)
		require.Len(t, res.ScrapJobs, 1)
		require.Equal(t, reqs[1].Site, res.ScrapJobs[0].Site)
		require.Equal(t, reqs[1].PostingId, res.ScrapJobs[0].PostingId)
	})
}

func newAddScrapJobRequest(userId string, num int) *restapi_grpc.AddScrapJobRequest {
	attachN := func(str string, n int) string {
		return fmt.Sprintf("%s%d", str, n)
	}
	return &restapi_grpc.AddScrapJobRequest{
		UserId:    userId,
		Site:      attachN("testSite", num),
		PostingId: attachN("testPostingId", num),
	}
}
