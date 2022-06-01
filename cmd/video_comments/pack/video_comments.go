package pack

import (
	"douyin/v1/cmd/video_comments/dal/db"
	"douyin/v1/kitex_gen/video_comments"
)

func Comment(m *db.Video_Comments) *video_comments.Video_Comments {
	if m == nil {
		return nil
	}
	return &video_comments.Video_Comments{
		CommentId:     m.Comment_ID,
		CommentUserId: m.Comment_User_ID,
		Content:       m.Content,
		CreateDate:    m.Create_Date,
	}
}
func Comments(ms []*db.Video_Comments) []*video_comments.Video_Comments {
	comments := make([]*video_comments.Video_Comments, 0)
	for _, m := range ms {
		if n := Comment(m); n != nil {
			comments = append(comments, n)
		}
	}
	return comments
}
