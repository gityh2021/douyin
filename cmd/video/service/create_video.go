package service

import (
	"context"
	"douyin/v1/cmd/video/dal/db"
	"douyin/v1/kitex_gen/video"
)

type CreateVideoService struct {
	ctx context.Context
}

func NewCreateVideoService(ctx context.Context) *CreateVideoService {
	return &CreateVideoService{ctx: ctx}
}
func (s *CreateVideoService) CreateVideo(video *video.Video) error {
	return db.MPublishVideo(s.ctx, &db.Video{
		AuthorID:      video.AuthorId,
		PlayURL:       video.PlayUrl,
		CoverURL:      video.CoverUrl,
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    video.IsFavorite,
		Title:         video.Title,
	})
}
