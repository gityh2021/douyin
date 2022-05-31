package pack

import (
	"douyin/v1/cmd/comment_action/dal/db"
	"douyin/v1/kitex_gen/video_comments"
)

func Comment(m *db.Comment_Action) *video_comments.Video_Comments {
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
