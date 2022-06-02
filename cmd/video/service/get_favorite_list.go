package service

import (
	"context"
	"douyin/v1/cmd/video/dal/db"
	"douyin/v1/cmd/video/pack"
	"douyin/v1/kitex_gen/video"
)

type QueryFavoriteService struct {
	ctx context.Context
}

func NewQueryFavoriteService(ctx context.Context) *QueryFavoriteService {
	return &QueryFavoriteService{ctx: ctx}
}

func (s *QueryFavoriteService) GetFavoriteListByUser(userId int64) ([]*video.Video, error) {
	videos, err := db.MGetFavoriteList(s.ctx, userId)
	if err != nil {
		return nil, err
	}
	return pack.Videos(videos), nil
}
