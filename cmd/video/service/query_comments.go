package service

import (
	"douyin/v1/cmd/video/dal/db"
	"douyin/v1/cmd/video/pack"
	"douyin/v1/kitex_gen/video"
)

func (s *CommentService) QueryCommentByVideoId(videoId int64) ([]*video.Comment, error) {
	comments, err := db.MGetCommentsByVideoId(s.ctx, videoId)
	if err != nil {
		return nil, err
	}
	return pack.Comments(comments), nil
}
