package service

import (
	"context"
	"douyin/v1/cmd/video/dal/db"
	"douyin/v1/cmd/video/pack"
	"douyin/v1/kitex_gen/video"
)

type CommentService struct {
	ctx context.Context
}

// NewCommentService 创建视频服务
func NewCommentService(ctx context.Context) *CommentService {
	return &CommentService{ctx: ctx}
}

// PostComment 给一条视频添加评论
func (s *CommentService) PostComment(comment *video.Comment) (*video.Comment, error) {
	dbComment := db.Comment{
		UserId:  comment.UserId,
		VideoId: comment.VideoId,
		Content: comment.Content,
	}
	err := db.MPostCommentAction(s.ctx, &dbComment)
	if err != nil {
		return nil, err
	}
	return pack.Comment(&dbComment), nil
}

// DeleteComment 删除一条评论
func (s *CommentService) DeleteComment(commentId int64) error {
	return db.MDeleteComment(s.ctx, commentId)
}
