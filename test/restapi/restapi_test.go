package restapi

import (
	"context"
	"fmt"
	"testing"

	"github.com/jae2274/careerhub-userinfo-service/careerhub/userinfo_service/restapi/restapi_grpc"
	"github.com/jae2274/careerhub-userinfo-service/test/tinit"
	"github.com/jae2274/goutils/ptr"
	"github.com/stretchr/testify/require"
)

func TestScrapJobGrpc(t *testing.T) {
	tinit.InitDB(t)
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

		IsExisted, err := client.RemoveScrapJob(ctx, &restapi_grpc.RemoveScrapJobRequest{
			UserId:    userId,
			Site:      reqs[0].Site,
			PostingId: reqs[0].PostingId,
		})
		require.NoError(t, err)
		require.True(t, IsExisted.IsExisted)

		res, err := client.GetScrapJobs(ctx, &restapi_grpc.GetScrapJobsRequest{UserId: userId})

		require.NoError(t, err)
		require.Len(t, res.ScrapJobs, 1)
		require.Equal(t, reqs[1].Site, res.ScrapJobs[0].Site)
		require.Equal(t, reqs[1].PostingId, res.ScrapJobs[0].PostingId)
	})

	t.Run("return isExisted false when remove non-existed", func(t *testing.T) {
		tinit.InitDB(t)
		ctx := context.Background()
		client := tinit.InitScrapJobGrpcClient(t)
		IsExisted, err := client.RemoveScrapJob(ctx, &restapi_grpc.RemoveScrapJobRequest{
			UserId:    "testUserId",
			Site:      "testSite",
			PostingId: "testPostingId",
		})
		require.NoError(t, err)
		require.False(t, IsExisted.IsExisted)
	})

	t.Run("return empty scrapJob cause did not scrap", func(t *testing.T) {
		tinit.InitDB(t)
		ctx := context.Background()
		client := tinit.InitScrapJobGrpcClient(t)

		jobPostingIds := []*restapi_grpc.JobPostingId{
			{Site: "testSite", PostingId: "testPostingId1"},
			{Site: "testSite", PostingId: "testPostingId2"},
			{Site: "testSite", PostingId: "testPostingId3"},
		}
		res, err := client.GetScrapJobsById(ctx, &restapi_grpc.GetScrapJobsByIdRequest{
			UserId:        "testUserId",
			JobPostingIds: jobPostingIds,
		})

		require.NoError(t, err)
		require.Empty(t, res.ScrapJobs)
	})

	t.Run("return scrapJob cause scrap", func(t *testing.T) {
		tinit.InitDB(t)
		ctx := context.Background()
		client := tinit.InitScrapJobGrpcClient(t)

		userId := "testUserId"
		jobPostingIds := []*restapi_grpc.JobPostingId{
			{Site: "testSite1", PostingId: "testPostingId1"},
			{Site: "testSite2", PostingId: "testPostingId2"},
			{Site: "testSite3", PostingId: "testPostingId3"},
		}

		for _, jobPostingId := range jobPostingIds[:2] {
			_, err := client.AddScrapJob(ctx, &restapi_grpc.AddScrapJobRequest{
				UserId:    userId,
				Site:      jobPostingId.Site,
				PostingId: jobPostingId.PostingId,
			})
			require.NoError(t, err)

			_, err = client.AddTag(ctx, &restapi_grpc.AddTagRequest{
				UserId:    userId,
				Site:      jobPostingId.Site,
				PostingId: jobPostingId.PostingId,
				Tag:       "testTag",
			})
			require.NoError(t, err)
		}
		_, err := client.AddScrapJob(ctx, &restapi_grpc.AddScrapJobRequest{
			UserId:    userId,
			Site:      "otherSite",
			PostingId: "otherPostingId",
		})
		require.NoError(t, err)
		_, err = client.AddScrapJob(ctx, &restapi_grpc.AddScrapJobRequest{
			UserId:    "otherUserId",
			Site:      "otherSite",
			PostingId: "otherPostingId",
		})
		require.NoError(t, err)

		res, err := client.GetScrapJobsById(ctx, &restapi_grpc.GetScrapJobsByIdRequest{
			UserId:        userId,
			JobPostingIds: jobPostingIds,
		})

		require.NoError(t, err)
		require.Len(t, res.ScrapJobs, 2)
		for i, jobPostingId := range jobPostingIds[:2] {
			require.Equal(t, jobPostingId.Site, res.ScrapJobs[i].Site)
			require.Equal(t, jobPostingId.PostingId, res.ScrapJobs[i].PostingId)
			require.Len(t, res.ScrapJobs[i].Tags, 1)
			require.Equal(t, "testTag", res.ScrapJobs[i].Tags[0])
		}
	})
}

func TestTags(t *testing.T) {
	tinit.InitDB(t)
	cancelFunc := tinit.RunTestApp(t)
	defer cancelFunc()

	t.Run("return empty", func(t *testing.T) {
		tinit.InitDB(t)
		ctx := context.Background()
		client := tinit.InitScrapJobGrpcClient(t)

		userId := "testUserId"
		_, err := client.AddScrapJob(ctx, newAddScrapJobRequest(userId, 1))
		require.NoError(t, err)

		res, err := client.GetScrapJobs(ctx, &restapi_grpc.GetScrapJobsRequest{UserId: userId})
		require.NoError(t, err)

		require.Len(t, res.ScrapJobs, 1)
		require.Empty(t, res.ScrapJobs[0].Tags)
	})

	t.Run("return one tag", func(t *testing.T) {
		tinit.InitDB(t)
		ctx := context.Background()
		client := tinit.InitScrapJobGrpcClient(t)

		userId := "testUserId"
		tag := "testTag"
		req := newAddScrapJobRequest(userId, 1)
		_, err := client.AddScrapJob(ctx, req)
		require.NoError(t, err)
		isExisted, err := client.AddTag(ctx, &restapi_grpc.AddTagRequest{
			UserId:    req.UserId,
			Site:      req.Site,
			PostingId: req.PostingId,
			Tag:       tag,
		})
		require.NoError(t, err)
		require.True(t, isExisted.IsExisted)

		res, err := client.GetScrapJobs(ctx, &restapi_grpc.GetScrapJobsRequest{UserId: userId})
		require.NoError(t, err)

		require.Len(t, res.ScrapJobs, 1)
		require.Len(t, res.ScrapJobs[0].Tags, 1)
		require.Equal(t, tag, res.ScrapJobs[0].Tags[0])
	})

	t.Run("return tags", func(t *testing.T) {
		tinit.InitDB(t)
		ctx := context.Background()
		client := tinit.InitScrapJobGrpcClient(t)

		userId := "testUserId"
		tags := []string{"testTag1", "testTag2"}
		req := newAddScrapJobRequest(userId, 1)
		_, err := client.AddScrapJob(ctx, req)
		require.NoError(t, err)

		for _, tag := range tags {
			isExisted, err := client.AddTag(ctx, &restapi_grpc.AddTagRequest{
				UserId:    req.UserId,
				Site:      req.Site,
				PostingId: req.PostingId,
				Tag:       tag,
			})
			require.NoError(t, err)
			require.True(t, isExisted.IsExisted)
		}

		res, err := client.GetScrapJobs(ctx, &restapi_grpc.GetScrapJobsRequest{UserId: userId})
		require.NoError(t, err)

		require.Len(t, res.ScrapJobs, 1)
		require.Len(t, res.ScrapJobs[0].Tags, len(tags))
		for i, tag := range tags {
			require.Equal(t, tag, res.ScrapJobs[0].Tags[i])
		}
	})

	t.Run("return unique tag though add duplicate tags", func(t *testing.T) {
		tinit.InitDB(t)
		ctx := context.Background()
		client := tinit.InitScrapJobGrpcClient(t)

		userId := "testUserId"
		tag := "testTag"
		req := newAddScrapJobRequest(userId, 1)
		_, err := client.AddScrapJob(ctx, req)
		require.NoError(t, err)

		for i := 0; i < 2; i++ {
			isExisted, err := client.AddTag(ctx, &restapi_grpc.AddTagRequest{
				UserId:    req.UserId,
				Site:      req.Site,
				PostingId: req.PostingId,
				Tag:       tag,
			})
			require.NoError(t, err)
			require.True(t, isExisted.IsExisted)
		}

		res, err := client.GetScrapJobs(ctx, &restapi_grpc.GetScrapJobsRequest{UserId: userId})
		require.NoError(t, err)

		require.Len(t, res.ScrapJobs, 1)
		require.Len(t, res.ScrapJobs[0].Tags, 1)
		require.Equal(t, tag, res.ScrapJobs[0].Tags[0])
	})

	t.Run("return tag without removed", func(t *testing.T) {
		tinit.InitDB(t)
		ctx := context.Background()
		client := tinit.InitScrapJobGrpcClient(t)

		userId := "testUserId"
		tags := []string{"testTag1", "testTag2"}
		req := newAddScrapJobRequest(userId, 1)
		_, err := client.AddScrapJob(ctx, req)
		require.NoError(t, err)

		for _, tag := range tags {
			isExisted, err := client.AddTag(ctx, &restapi_grpc.AddTagRequest{
				UserId:    req.UserId,
				Site:      req.Site,
				PostingId: req.PostingId,
				Tag:       tag,
			})
			require.NoError(t, err)
			require.True(t, isExisted.IsExisted)
		}

		isExisted, err := client.RemoveTag(ctx, &restapi_grpc.RemoveTagRequest{
			UserId:    req.UserId,
			Site:      req.Site,
			PostingId: req.PostingId,
			Tag:       tags[0],
		})
		require.NoError(t, err)
		require.True(t, isExisted.IsExisted)

		res, err := client.GetScrapJobs(ctx, &restapi_grpc.GetScrapJobsRequest{UserId: userId})
		require.NoError(t, err)

		require.Len(t, res.ScrapJobs, 1)
		require.Len(t, res.ScrapJobs[0].Tags, 1)
		require.Equal(t, tags[1], res.ScrapJobs[0].Tags[0])
	})

	t.Run("return empty scrapTags ", func(t *testing.T) {
		tinit.InitDB(t)
		ctx := context.Background()
		client := tinit.InitScrapJobGrpcClient(t)

		userId := "testUserId"
		res, err := client.GetScrapTags(ctx, &restapi_grpc.GetScrapTagsRequest{UserId: userId})
		require.NoError(t, err)

		require.Empty(t, res.Tags)
	})

	t.Run("return scrapTags from one scrapJob", func(t *testing.T) {
		tinit.InitDB(t)
		ctx := context.Background()
		client := tinit.InitScrapJobGrpcClient(t)

		userId := "testUserId"
		tags := []string{"testTag1", "testTag2"}
		req := newAddScrapJobRequest(userId, 1)
		_, err := client.AddScrapJob(ctx, req)
		require.NoError(t, err)

		for _, tag := range tags {
			isExisted, err := client.AddTag(ctx, &restapi_grpc.AddTagRequest{
				UserId:    req.UserId,
				Site:      req.Site,
				PostingId: req.PostingId,
				Tag:       tag,
			})
			require.NoError(t, err)
			require.True(t, isExisted.IsExisted)
		}

		res, err := client.GetScrapTags(ctx, &restapi_grpc.GetScrapTagsRequest{UserId: userId})
		require.NoError(t, err)
		require.Len(t, res.Tags, len(tags))
		require.Equal(t, tags, res.Tags)
	})

	t.Run("return scrapTags from multiple scrapJobs", func(t *testing.T) {
		tinit.InitDB(t)
		ctx := context.Background()
		client := tinit.InitScrapJobGrpcClient(t)

		userId := "testUserId"
		tags1 := []string{"testTag1", "testTag2"}
		tags2 := []string{"testTag3", "testTag4"}
		req1 := newAddScrapJobRequest(userId, 1)
		req2 := newAddScrapJobRequest(userId, 2)
		_, err := client.AddScrapJob(ctx, req1)
		require.NoError(t, err)
		_, err = client.AddScrapJob(ctx, req2)
		require.NoError(t, err)

		for _, tag := range tags1 {
			isExisted, err := client.AddTag(ctx, &restapi_grpc.AddTagRequest{
				UserId:    req1.UserId,
				Site:      req1.Site,
				PostingId: req1.PostingId,
				Tag:       tag,
			})
			require.NoError(t, err)
			require.True(t, isExisted.IsExisted)
		}

		for _, tag := range tags2 {
			isExisted, err := client.AddTag(ctx, &restapi_grpc.AddTagRequest{
				UserId:    req2.UserId,
				Site:      req2.Site,
				PostingId: req2.PostingId,
				Tag:       tag,
			})
			require.NoError(t, err)
			require.True(t, isExisted.IsExisted)
		}

		res, err := client.GetScrapTags(ctx, &restapi_grpc.GetScrapTagsRequest{UserId: userId})
		require.NoError(t, err)
		require.Len(t, res.Tags, len(tags1)+len(tags2))
		require.Contains(t, res.Tags, tags1[0])
		require.Contains(t, res.Tags, tags1[1])
		require.Contains(t, res.Tags, tags2[0])
		require.Contains(t, res.Tags, tags2[1])
	})

	t.Run("return unique scrapTags though have duplicate tags", func(t *testing.T) {
		tinit.InitDB(t)
		ctx := context.Background()
		client := tinit.InitScrapJobGrpcClient(t)

		userId := "testUserId"
		tags1 := []string{"testTag1", "testTag2", "testTag3"}
		tags2 := []string{"testTag1", "testTag2", "testTag4"}
		req1 := newAddScrapJobRequest(userId, 1)
		req2 := newAddScrapJobRequest(userId, 2)
		_, err := client.AddScrapJob(ctx, req1)
		require.NoError(t, err)
		_, err = client.AddScrapJob(ctx, req2)
		require.NoError(t, err)

		for _, tag := range tags1 {
			isExisted, err := client.AddTag(ctx, &restapi_grpc.AddTagRequest{
				UserId:    req1.UserId,
				Site:      req1.Site,
				PostingId: req1.PostingId,
				Tag:       tag,
			})
			require.NoError(t, err)
			require.True(t, isExisted.IsExisted)
		}

		for _, tag := range tags2 {
			isExisted, err := client.AddTag(ctx, &restapi_grpc.AddTagRequest{
				UserId:    req2.UserId,
				Site:      req2.Site,
				PostingId: req2.PostingId,
				Tag:       tag,
			})
			require.NoError(t, err)
			require.True(t, isExisted.IsExisted)
		}

		res, err := client.GetScrapTags(ctx, &restapi_grpc.GetScrapTagsRequest{UserId: userId})
		require.NoError(t, err)
		require.Len(t, res.Tags, 4)
		require.Contains(t, res.Tags, tags1[0])
		require.Contains(t, res.Tags, tags1[1])
		require.Contains(t, res.Tags, tags1[2])
		require.Contains(t, res.Tags, tags2[2])
	})

	t.Run("return scrapTags without removed tag", func(t *testing.T) {
		tinit.InitDB(t)
		ctx := context.Background()
		client := tinit.InitScrapJobGrpcClient(t)

		userId := "testUserId"
		tags := []string{"testTag1", "testTag2"}
		req := newAddScrapJobRequest(userId, 1)
		_, err := client.AddScrapJob(ctx, req)
		require.NoError(t, err)

		for _, tag := range tags {
			isExisted, err := client.AddTag(ctx, &restapi_grpc.AddTagRequest{
				UserId:    req.UserId,
				Site:      req.Site,
				PostingId: req.PostingId,
				Tag:       tag,
			})
			require.NoError(t, err)
			require.True(t, isExisted.IsExisted)
		}

		isExisted, err := client.RemoveTag(ctx, &restapi_grpc.RemoveTagRequest{
			UserId:    req.UserId,
			Site:      req.Site,
			PostingId: req.PostingId,
			Tag:       tags[0],
		})
		require.NoError(t, err)
		require.True(t, isExisted.IsExisted)

		res, err := client.GetScrapTags(ctx, &restapi_grpc.GetScrapTagsRequest{UserId: userId})
		require.NoError(t, err)
		require.Len(t, res.Tags, 1)
		require.Equal(t, tags[1], res.Tags[0])
	})

	t.Run("return scrapTags with different userId", func(t *testing.T) {
		tinit.InitDB(t)
		ctx := context.Background()
		client := tinit.InitScrapJobGrpcClient(t)

		userId1 := "testUserId1"
		userId2 := "testUserId2"
		tags1 := []string{"testTag1", "testTag2"}
		tags2 := []string{"testTag3", "testTag4"}
		req1 := newAddScrapJobRequest(userId1, 1)
		req2 := newAddScrapJobRequest(userId2, 1)
		_, err := client.AddScrapJob(ctx, req1)
		require.NoError(t, err)
		_, err = client.AddScrapJob(ctx, req2)
		require.NoError(t, err)

		for _, tag := range tags1 {
			isExisted, err := client.AddTag(ctx, &restapi_grpc.AddTagRequest{
				UserId:    req1.UserId,
				Site:      req1.Site,
				PostingId: req1.PostingId,
				Tag:       tag,
			})
			require.NoError(t, err)
			require.True(t, isExisted.IsExisted)
		}

		for _, tag := range tags2 {
			isExisted, err := client.AddTag(ctx, &restapi_grpc.AddTagRequest{
				UserId:    req2.UserId,
				Site:      req2.Site,
				PostingId: req2.PostingId,
				Tag:       tag,
			})
			require.NoError(t, err)
			require.True(t, isExisted.IsExisted)
		}

		res2, err := client.GetScrapTags(ctx, &restapi_grpc.GetScrapTagsRequest{UserId: userId2})
		require.NoError(t, err)
		require.Len(t, res2.Tags, len(tags2))
		require.Contains(t, res2.Tags, tags2[0])
		require.Contains(t, res2.Tags, tags2[1])
	})

	t.Run("return empty scrapJob cause empty tag", func(t *testing.T) {
		tinit.InitDB(t)
		ctx := context.Background()
		client := tinit.InitScrapJobGrpcClient(t)

		userId := "testUserId"
		tag := "testTag"
		req := newAddScrapJobRequest(userId, 1)
		_, err := client.AddScrapJob(ctx, req)
		require.NoError(t, err)

		res, err := client.GetScrapJobs(ctx, &restapi_grpc.GetScrapJobsRequest{UserId: userId, Tag: ptr.P(tag)})
		require.NoError(t, err)
		require.Empty(t, res.ScrapJobs)
	})

	t.Run("return empty scrapJob cause different tag", func(t *testing.T) {
		tinit.InitDB(t)
		ctx := context.Background()
		client := tinit.InitScrapJobGrpcClient(t)

		userId := "testUserId"
		tag := "testTag"
		diffTag := "diffTag"
		req := newAddScrapJobRequest(userId, 1)
		_, err := client.AddScrapJob(ctx, req)
		require.NoError(t, err)

		isExisted, err := client.AddTag(ctx, &restapi_grpc.AddTagRequest{
			UserId:    req.UserId,
			Site:      req.Site,
			PostingId: req.PostingId,
			Tag:       diffTag,
		})
		require.NoError(t, err)
		require.True(t, isExisted.IsExisted)

		res, err := client.GetScrapJobs(ctx, &restapi_grpc.GetScrapJobsRequest{UserId: userId, Tag: ptr.P(tag)})
		require.NoError(t, err)
		require.Empty(t, res.ScrapJobs)
	})

	t.Run("return scrapJob with tag", func(t *testing.T) {
		tinit.InitDB(t)
		ctx := context.Background()
		client := tinit.InitScrapJobGrpcClient(t)

		userId := "testUserId"
		tag := "testTag"
		req := newAddScrapJobRequest(userId, 1)
		_, err := client.AddScrapJob(ctx, req)
		require.NoError(t, err)

		isExisted, err := client.AddTag(ctx, &restapi_grpc.AddTagRequest{
			UserId:    req.UserId,
			Site:      req.Site,
			PostingId: req.PostingId,
			Tag:       tag,
		})
		require.NoError(t, err)
		require.True(t, isExisted.IsExisted)

		res, err := client.GetScrapJobs(ctx, &restapi_grpc.GetScrapJobsRequest{UserId: userId, Tag: ptr.P(tag)})
		require.NoError(t, err)
		require.Len(t, res.ScrapJobs, 1)
		require.Equal(t, req.Site, res.ScrapJobs[0].Site)
		require.Equal(t, req.PostingId, res.ScrapJobs[0].PostingId)
	})

	t.Run("return scrapJobs with tag from multiple scrapJobs", func(t *testing.T) {
		tinit.InitDB(t)
		ctx := context.Background()
		client := tinit.InitScrapJobGrpcClient(t)

		userId := "testUserId"
		tag := "testTag1"

		tags1 := []string{tag, "testTag2"}
		tags2 := []string{tag, "testTag3"}
		tags3 := []string{"testTag3", "testTag5"}
		req1 := newAddScrapJobRequest(userId, 1)
		req2 := newAddScrapJobRequest(userId, 2)
		req3 := newAddScrapJobRequest(userId, 3)
		reqs := []*restapi_grpc.AddScrapJobRequest{req1, req2, req3}
		tagsList := [][]string{tags1, tags2, tags3}

		for i, req := range reqs {
			_, err := client.AddScrapJob(ctx, req)
			require.NoError(t, err)

			for _, tag := range tagsList[i] {
				isExisted, err := client.AddTag(ctx, &restapi_grpc.AddTagRequest{
					UserId:    req.UserId,
					Site:      req.Site,
					PostingId: req.PostingId,
					Tag:       tag,
				})
				require.NoError(t, err)
				require.True(t, isExisted.IsExisted)
			}
		}

		res, err := client.GetScrapJobs(ctx, &restapi_grpc.GetScrapJobsRequest{UserId: userId, Tag: ptr.P(tag)})
		require.NoError(t, err)
		require.Len(t, res.ScrapJobs, 2)
		for i, req := range reqs[:2] {
			require.Equal(t, req.Site, res.ScrapJobs[i].Site)
			require.Equal(t, req.PostingId, res.ScrapJobs[i].PostingId)
		}
	})

	t.Run("return isExisted false when add tag to non-existed scrapJob", func(t *testing.T) {
		tinit.InitDB(t)
		ctx := context.Background()
		client := tinit.InitScrapJobGrpcClient(t)
		IsExisted, err := client.AddTag(ctx, &restapi_grpc.AddTagRequest{
			UserId:    "testUserId",
			Site:      "testSite",
			PostingId: "testPostingId",
			Tag:       "testTag",
		})
		require.NoError(t, err)
		require.False(t, IsExisted.IsExisted)
	})

	t.Run("return isExisted false when remove tag to non-existed scrapJob", func(t *testing.T) {
		tinit.InitDB(t)
		ctx := context.Background()
		client := tinit.InitScrapJobGrpcClient(t)
		IsExisted, err := client.RemoveTag(ctx, &restapi_grpc.RemoveTagRequest{
			UserId:    "testUserId",
			Site:      "testSite",
			PostingId: "testPostingId",
			Tag:       "testTag",
		})
		require.NoError(t, err)
		require.False(t, IsExisted.IsExisted)
	})

	t.Run("return isExisted true when remove non-existed tag to scrapJob", func(t *testing.T) { //스크랩한 채용공고 대상으로 태그를 삭제할 때, 태그가 존재하지 않아도 성공해야 합니다.
		tinit.InitDB(t)
		ctx := context.Background()
		client := tinit.InitScrapJobGrpcClient(t)

		userId := "testUserId"
		tag := "testTag"
		req := newAddScrapJobRequest(userId, 1)
		_, err := client.AddScrapJob(ctx, req)
		require.NoError(t, err)

		isExisted, err := client.RemoveTag(ctx, &restapi_grpc.RemoveTagRequest{
			UserId:    req.UserId,
			Site:      req.Site,
			PostingId: req.PostingId,
			Tag:       tag,
		})
		require.NoError(t, err)
		require.True(t, isExisted.IsExisted)
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
