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
