package rpc

import (
	"context"
	"douyin/v1/kitex_gen/video"
	"douyin/v1/kitex_gen/video/videoservice"
	"douyin/v1/pkg/constants"
	"douyin/v1/pkg/errno"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var videoClient videoservice.Client

func initVideoRpc() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}
	c, err := videoservice.NewClient(
		"video",
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(3*time.Second),              // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond),    // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithResolver(r),                            // resolver
	)
	if err != nil {
		panic(err)
	}
	videoClient = c
}

func GetPublishVideoList(ctx context.Context, userId int64) ([]*video.Video, error) {
	resp, err := videoClient.GetPublishListByUser(ctx, userId)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return resp.VideoList, nil
}

func GetVideosFeed(ctx context.Context, lastTime int64, userId int64) ([]*video.Video, int64, error) {
	resp, err := videoClient.GetVideosByLastTime(ctx, lastTime, userId)
	if err != nil {
		return nil, time.Now().Unix(), err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, resp.NextTime, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return resp.VideoList, resp.NextTime, nil
}

func CreateVideo(ctx context.Context, video *video.Video) error {
	resp, err := videoClient.PublishVideo(ctx, video)
	if err != nil {
		return err
	}
	if resp.StatusCode != 0 {
		return errno.NewErrNo(resp.StatusCode, resp.StatusMsg)
	}
	return nil
}

func GetFavoriteList(ctx context.Context, userId int64) ([]*video.Video, error) {
	resp, err := videoClient.GetFavoriteListBYUser(ctx, userId)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return resp.VideoList, nil
}

func FavoriteByUser(ctx context.Context, request *video.FavoriteActionRequest) (*video.BaseResp, error) {
	resp, err := videoClient.FavoriteByUser(ctx, request)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.StatusCode, resp.StatusMsg)
	}
	return resp, nil
}

func PostComment(ctx context.Context, request *video.CommentActionRequest) (*video.CommentActionResponse, error) {
	resp, err := videoClient.PostComment(ctx, request)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return resp, nil
}

func GetCommentsByVideoId(ctx context.Context, videoId int64) ([]*video.Comment, error) {
	resp, err := videoClient.GetCommentListByVideo(ctx, videoId)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return resp.CommentList, nil
}
