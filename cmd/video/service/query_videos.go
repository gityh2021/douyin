package service

import (
	"context"
	"douyin/v1/cmd/video/dal/db"
	"douyin/v1/cmd/video/pack"
	"douyin/v1/kitex_gen/video"
	"time"
)

type QueryVideoService struct {
	ctx context.Context
}

func NewQueryVideoService(ctx context.Context) *QueryVideoService {
	return &QueryVideoService{ctx: ctx}
}

func (s *QueryVideoService) GetPublishList(userId int64) ([]*video.Video, error) {
	videos, err := db.MGetVideosByUserID(s.ctx, userId)
	if err != nil {
		return nil, err
	}
	return pack.Videos(videos), nil
}

func (s *QueryVideoService) GetVideoFeed(lastTime int64) ([]*video.Video, time.Time, error) {
	videos, err := db.MGetVideosByTime(s.ctx, lastTime)
	if err != nil {
		return nil, time.Now(), err
	}
	if len(videos) == 0 {
		return pack.Videos(videos), time.Now(), nil
	}
	return pack.Videos(videos), videos[0].CreatedAt, nil
}
