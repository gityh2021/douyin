package db

import (
	"context"
	"gorm.io/gorm"
)

type Favorite struct {
	gorm.Model
	ID      int64 `gorm:"primarykey" json:"id"`
	UserId  int64 `gorm:"not null" json:"user_id"`
	VideoId int64 `gorm:"not null" json:"video_id"`
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
	if err := DB.WithContext(ctx).Create(&favorite).Error; err != nil {
		return err
	}
	return DB.WithContext(ctx).Model(&Video{}).Where("id = ?", videoId).Update("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error
}

func MCancelFavoriteAction(ctx context.Context, userId int64, videoId int64) error {
	if err := DB.WithContext(ctx).Where("user_id = ? and video_id = ?", userId, videoId).Delete(&Favorite{}).Error; err != nil {
		return err
	}
	return DB.WithContext(ctx).Model(&Video{}).Where("id = ?", videoId).Update("favorite_count", gorm.Expr("favorite_count - ?", 1)).Error
}

// MGetFavoriteIds 返回userId点赞过的videoIds内的视频
func MGetFavoriteIds(ctx context.Context, videoIds []*int64, userId int64) ([]*int64, error) {
	res := make([]*int64, 0)
	if err := DB.WithContext(ctx).Model(&Favorite{}).Select("video_id").Where("user_id = ? and video_id in (?)", userId, videoIds).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
