package db

import (
	"context"
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

func MGetVideos(ctx context.Context, userId int64) ([]*Video, error) {
	res := make([]*Video, 0)
	if err := DB.WithContext(ctx).Where("author_id = ?", userId).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
