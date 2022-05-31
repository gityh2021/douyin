package main

import (
	"context"
	"douyin/v1/cmd/video_comments/pack"
	"douyin/v1/cmd/video_comments/service"
	"douyin/v1/kitex_gen/video_comments"
	"douyin/v1/pkg/errno"
)

// VideoCommentsServiceImpl implements the last service interface defined in the IDL.
type VideoCommentsServiceImpl struct{}

// GetPublishVideoCommentByVideo implements the VideoCommentsServiceImpl interface.
func (s *VideoCommentsServiceImpl) GetPublishVideoCommentByVideo(ctx context.Context, videoId int64) (resp *video_comments.PublishVideoCommentResponse, err error) {
	response := new(video_comments.PublishVideoCommentResponse)
	if videoId < 0 {
		response.SetVcresp(pack.BuildVideoCommentResp(errno.ParamErr))
		return response, nil
	}
	comments, err := service.NewQueryVideoCommentService(ctx).GetPublishList(videoId)
	if err != nil {
		response.SetVcresp(pack.BuildVideoCommentResp(err))
		return response, nil
	}
	response.SetVcresp(pack.BuildVideoCommentResp(errno.Success))
	response.SetVideoComments(comments)
	return response, nil
}
