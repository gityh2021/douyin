package db

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	ID            int64  `gorm:"primarykey" json:"user_name"`
	AuthorID      int64  `gorm:"not null" json:"author_id"`
	PlayURL       string `gorm:"not null" json:"play_url"`
	CoverURL      string `gorm:"not null" json:"cover_url"`
	FavoriteCount int64  `gorm:"default:0" json:"favorite_count"`
	CommentCount  int64  `gorm:"default:0" json:"comment_count"`
	IsFavorite    bool   `json:"is_favorite"`
	Title         string `json:"title"`
}

// MGetVideosByUserID 根据Userid返回视频
func MGetVideosByUserID(ctx context.Context, userId int64) ([]*Video, error) {
	res := make([]*Video, 0)
	if err := DB.WithContext(ctx).Where("author_id = ?", userId).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// MGetVideosByTime 根据时间返回视频
func MGetVideosByTime(ctx context.Context, lastTime int64) ([]*Video, error) {
	res := make([]*Video, 0)
	if err := DB.WithContext(ctx).Where("created_at < ?", time.Unix(lastTime, 0).Format("2006-01-02 15:04:05")).Order("created_at desc").Limit(30).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// MPublishVideo 发布视频
func MPublishVideo(ctx context.Context, video *Video) error {
	return DB.WithContext(ctx).Create(&video).Error
}
