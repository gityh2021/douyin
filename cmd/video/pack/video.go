package pack

import (
	"douyin/v1/cmd/video/dal/db"
	"douyin/v1/kitex_gen/video"
)

func Video(m *db.Video) *video.Video {
	if m == nil {
		return nil
	}
	return &video.Video{
		Id:            m.ID,
		AuthorId:      m.AuthorID,
		PlayUrl:       m.PlayURL,
		CoverUrl:      m.CoverURL,
		FavoriteCount: m.FavoriteCount,
		CommentCount:  m.CommentCount,
		IsFavorite:    m.IsFavorite,
		Title:         m.Title,
	}
}

func Videos(ms []*db.Video) []*video.Video {
	notes := make([]*video.Video, 0)
	for _, m := range ms {
		if n := Video(m); n != nil {
			notes = append(notes, n)
		}
	}
	return notes
}

func Comment(m *db.Comment) *video.Comment {
	if m == nil {
		return nil
	}
	return &video.Comment{
		Id:         m.ID,
		UserId:     m.UserId,
		VideoId:    m.VideoId,
		Content:    m.Content,
		CreateDate: m.CreatedAt.Format("01-02"),
	}
}

func Comments(ms []*db.Comment) []*video.Comment {
	notes := make([]*video.Comment, 0)
	for _, m := range ms {
		if n := Comment(m); n != nil {
			notes = append(notes, n)
		}
	}
	return notes
}
