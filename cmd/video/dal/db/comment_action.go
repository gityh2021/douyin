package db

import (
	"context"
	"douyin/v1/cmd/video_comments/dal/db"
	"douyin/v1/kitex_gen/video_comments"
)

type Comment_Action struct {
	User_id      int64
	Token        string
	Video_id     int64
	Action_type  int64
	Comment_text string
	Comment_id   int64
	Create_date  string
}

// parse comment_action to comment
func Comment(m *Comment_Action) *video_comments.Video_Comments {
	if m == nil {
		return nil
	}
	return &video_comments.Video_Comments{
		CommentId:     m.Comment_id,
		CommentUserId: m.User_id,
		Content:       m.Comment_text,
		CreateDate:    m.Create_date,
	}
}

// depend on action_type to do something
func VideoCommentAction(ctx context.Context, comment_action *Comment_Action) error {
	action_type := comment_action.Action_type
	switch action_type {
	case 1:
		return CreateVideoComment(ctx, comment_action)
	case 2:
		return DeleteVideoComment(ctx, comment_action)
	default:
		panic(action_type)
	}
}

// create video comment
func CreateVideoComment(ctx context.Context, comment_action *Comment_Action) error {
	if err := DB.WithContext(ctx).Create(Comment(comment_action)).Error; err != nil {
		return err
	}
	return nil
}

// delete video comment
func DeleteVideoComment(ctx context.Context, comment_action *Comment_Action) error {
	return DB.WithContext(ctx).Where("comment_id = ? ", comment_action.Comment_id).Delete(&db.Video_Comments{}).Error
}
