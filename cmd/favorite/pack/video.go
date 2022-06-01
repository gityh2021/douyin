package pack

import (
	"douyin/v1/cmd/video/dal/db"
	"douyin/v1/kitex_gen/favorite"
)

func Video(m *db.Video) *favorite.Video {
	if m == nil {

		return nil
	}
	return &favorite.Video{
		Id:            m.ID,
		AuthorId:      m.AuthorID,
		PlayUrl:       m.PlayURL,
		CoverUrl:      m.PlayURL,
		FavoriteCount: m.FavoriteCount,
		CommentCount:  m.CommentCount,
		IsFavorite:    m.IsFavorite,
		Title:         m.Title,
	}
}
func Videos(ms []*db.Video) []*favorite.Video {
	notes := make([]*favorite.Video, 0)
	for _, m := range ms {
		if n := Video(m); n != nil {
			notes = append(notes, n)
		}
	}
	return notes
}
