package db

import (
	"context"
	"gorm.io/gorm"
)

type Favorite struct {
	gorm.Model
	ID      int64 `gorm:"primarykey" json:"id"`
	UserId  int64 `gorm:"default:0" json:"user_id"`
	VideoId int64 `gorm:"default:0" json:"video_id"`
}

func MGetFavoriteList(ctx context.Context, userId int64) ([]*Video, error) {
	videos := make([]*Video, 0)
	ids := make([]int64, 0)
	if err := DB.WithContext(ctx).Where("id in (?)",
		DB.WithContext(ctx).Table("favorites").Select("video_id").Where("user_id = ?", userId).Find(&ids)).Find(&videos).Error; err != nil {
		return nil, err
	}
	return videos, nil
}

func MPostFavoriteAction(ctx context.Context, userId int64, videoId int64) error {
	favorite := Favorite{
		UserId:  userId,
		VideoId: videoId,
	}
	return DB.WithContext(ctx).Create(&favorite).Error
}
