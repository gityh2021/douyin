package service

import (
	"context"
	"douyin/v1/cmd/video/dal/db"
	"douyin/v1/cmd/video/pack"
	"douyin/v1/kitex_gen/video"
)

type QueryVideoService struct {
	ctx context.Context
}

func NewQueryVideoService(ctx context.Context) *QueryVideoService{
	return &QueryVideoService{ctx: ctx}
}

func (s *QueryVideoService) GetPublishList(userId int64) ([]*video.Video, error) {
	videos, err := db.MGetVideos(s.ctx, userId)
	if err != nil{
		return nil, err
	}
	return pack.Videos(videos), nil
}
