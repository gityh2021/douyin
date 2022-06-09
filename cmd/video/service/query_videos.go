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

// NewQueryVideoService 创建视频查询服务
func NewQueryVideoService(ctx context.Context) *QueryVideoService {
	return &QueryVideoService{ctx: ctx}
}

// GetPublishList 获取用户的作品列表
func (s *QueryVideoService) GetPublishList(userId int64) ([]*video.Video, error) {
	videos, err := db.MGetVideosByUserID(s.ctx, userId)
	if err != nil {
		return nil, err
	}
	if len(videos) == 0 {
		return pack.Videos(videos), nil
	}
	newVideos, err := NewQueryFavoriteService(s.ctx).AppendFavorite(videos, userId)
	if err != nil {
		return nil, err
	}
	return pack.Videos(newVideos), nil
}

// GetVideoFeed 获取用户最新的视频列表
func (s *QueryVideoService) GetVideoFeed(lastTime int64, userId int64) ([]*video.Video, time.Time, error) {
	videos, err := db.MGetVideosByTime(s.ctx, lastTime)
	if err != nil {
		return nil, time.Now(), err
	}
	if len(videos) == 0 {
		return pack.Videos(videos), time.Now(), nil
	}
	newVideos, err := NewQueryFavoriteService(s.ctx).AppendFavorite(videos, userId)
	if err != nil {
		return nil, time.Now(), err
	}
	return pack.Videos(newVideos), videos[0].CreatedAt, nil
}
