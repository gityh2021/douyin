package main

import (
	"context"
	"douyin/v1/cmd/comment_action/pack"
	"douyin/v1/cmd/comment_action/service"
	"douyin/v1/kitex_gen/comment_action"
	"douyin/v1/pkg/errno"
)

// CommentActionServiceImpl implements the last service interface defined in the IDL.
type CommentActionServiceImpl struct{}

// PostCommentActionResponse implements the CommentActionServiceImpl interface.
func (s *CommentActionServiceImpl) PostCommentActionResponse(ctx context.Context, req *comment_action.Comment_Action) (resp *comment_action.CommentActionResponse, err error) {
	resp = new(comment_action.CommentActionResponse)

	if req.UserId <= 0 || req.CommentId <= 0 || len(req.CommentText) == 0 {
		resp.Caresp = pack.BuildCommentActionResp(errno.ParamErr)
		return resp, nil
	}

	err = service.NewCommentService(ctx).CommentService(req)
	if err != nil {
		resp.Caresp = pack.BuildCommentActionResp(err)
		return resp, nil
	}
	resp.Caresp = pack.BuildCommentActionResp(errno.Success)
	return resp, nil
}
