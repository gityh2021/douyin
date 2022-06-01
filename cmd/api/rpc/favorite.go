package rpc

import (
	"context"
	"douyin/v1/kitex_gen/favorite"
	"douyin/v1/kitex_gen/favorite/favoriteservice"
	"douyin/v1/pkg/errno"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	"time"
)

var favoriteClient favoriteservice.Client

func initFavoriteRpc() {
	r, err := etcd.NewEtcdResolver([]string{"139.224.195.12:2379"})
	if err != nil {
		panic(err)
	}

	c, err := favoriteservice.NewClient(
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
	favoriteClient = c
}

func GetFavoriteList(ctx context.Context, request *favorite.FavoriteListRequest) ([]*favorite.Video, error) {
	resp, err := favoriteClient.GetFavoriteListBYUser(ctx, request)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return resp.VideoList, nil
}

func FavoriteByUser(ctx context.Context, request *favorite.FavoriteActionRequest) (*favorite.BaseResponse, error) {
	resp, err := favoriteClient.FavoriteByUser(ctx, request)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.StatusCode, resp.StatusMessage)
	}
	return resp, nil
}
