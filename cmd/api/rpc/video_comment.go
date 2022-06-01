package rpc

import (
	"context"
	"douyin/v1/kitex_gen/video_comments"
	"douyin/v1/kitex_gen/comment_action"
	"douyin/v1/kitex_gen/video_comments/videocommentsservice"
	"douyin/v1/kitex_gen/comment_action/commentactionservice"
	"douyin/v1/pkg/errno"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	"time"
)
var videocommentsClient videocommentsservice.Client
var commentactionClient commentactionservice.Client

func initVideoCommentsRpc() {
	r, err := etcd.NewEtcdResolver([]string{"139.224.195.12:2379"})
	if err != nil {
		panic(err)
	}

	c, err := videocommentsservice.NewClient(
		"video_comments",
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(3*time.Second),              // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond),    // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithResolver(r),                            // resolver
	)
	if err != nil {
		panic(err)
	}
	videocommentsClient = c
}

func initCommentActionRpc() {
	r, err := etcd.NewEtcdResolver([]string{"139.224.195.12:2379"})
	if err != nil {
		panic(err)
	}

	c, err := commentactionservice.NewClient(
		"comment_action",
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(3*time.Second),              // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond),    // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithResolver(r),                            // resolver
	)
	if err != nil {
		panic(err)
	}
	commentactionClient = c
}

func GetVideoCommentsList(ctx context.Context,  videoId int64) ([]*video_comments.Video_Comments, error) {
	resp, err := videocommentsClient.GetPublishVideoCommentByVideo(ctx, videoId)
	if err != nil{
		return nil, err
	}
	if resp.Vcresp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.Vcresp.StatusCode, resp.Vcresp.StatusMessage)
	}
	return resp.VideoComments, nil
}

// Create or Delete Comment comment info
func CommentAction(ctx context.Context, req *comment_action.Comment_Action) error {
	resp, err := commentactionClient.PostCommentActionResponse(ctx, req)
	if err != nil {
		return err
	}
	if resp.Caresp.StatusCode != 0 {
		return errno.NewErrNo(resp.Caresp.StatusCode, resp.Caresp.StatusMessage)
	}
	return nil
}
