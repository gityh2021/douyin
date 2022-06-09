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

// NewQueryFavoriteService 创建查询视频点赞服务
func NewQueryFavoriteService(ctx context.Context) *QueryFavoriteService {
	return &QueryFavoriteService{ctx: ctx}
}

// GetFavoriteListByUser 获取用户点赞的视频列表
func (s *QueryFavoriteService) GetFavoriteListByUser(userId int64) ([]*video.Video, error) {
	videos, err := db.MGetFavoriteList(s.ctx, userId)
	if err != nil {
		return nil, err
	}
	return pack.Videos(videos), nil
}

// AppendFavorite 给一条视频点赞
func (s *QueryFavoriteService) AppendFavorite(ms []*db.Video, userId int64) ([]*db.Video, error) {
	if len(ms) == 0 {
		return ms, nil
	}
	videoIds := make([]*int64, 0)
	for _, m := range ms {
		videoIds = append(videoIds, &m.ID)
	}
	ids, err := db.MGetFavoriteIds(s.ctx, videoIds, userId)
	if err != nil {
		return ms, err
	}
	for _, m := range ms {
		for _, id := range ids {
			if m.ID == *id {
				m.IsFavorite = true
				break
			}
		}
	}
	return ms, nil
}
