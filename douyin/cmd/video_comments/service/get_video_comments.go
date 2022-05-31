package service

import (
	"context"
	"douyin/v1/cmd/video_comments/dal/db"
	"douyin/v1/cmd/video_comments/pack"
	"douyin/v1/kitex_gen/video_comments"
)

type QueryVideoCommentService struct {
	ctx context.Context
}

func NewQueryVideoCommentService(ctx context.Context) *QueryVideoCommentService {
	return &QueryVideoCommentService{ctx: ctx}
}

func (s *QueryVideoCommentService) GetPublishList(videoId int64) ([]*video_comments.Video_Comments, error) {
	comments, err := db.MGetVideosComments(s.ctx, videoId)
	if err != nil {
		return nil, err
	}
	return pack.Comments(comments), nil
}
