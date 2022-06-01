package rpc

import (
	"context"
	"douyin/v1/kitex_gen/video"
	"douyin/v1/kitex_gen/video/videoservice"
	"douyin/v1/pkg/errno"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	"time"
)
var videoClient videoservice.Client

func initVideoRpc() {
	r, err := etcd.NewEtcdResolver([]string{"139.224.195.12:2379"})
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

func GetPublishVideoList(ctx context.Context,  userId int64) ([]*video.Video, error) {
	resp, err := videoClient.GetPublishListByUser(ctx, userId)
	if err != nil{
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return resp.VideoList, nil
}