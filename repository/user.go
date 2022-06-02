package repository

import (
	"context"
	"douyin/v1/cmd/video/dal/db"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name          string `gorm:"not null" json:"name"`
	ID            uint64 `gorm:"cprimarykey"`
	Username      string `gorm:"column:username"`
	Password      string `gorm:"column:password"`
	FollowCount   int64  `gorm:"default:0" json:"follow_count"`
	FollowerCount int64  `gorm:"default:0" json:"follower_count"`
}

func MGetFavoriteList(ctx context.Context, userId int64) ([]*db.Video, error) {
	res := make([]*db.Video, 0)
	if err := DB.WithContext(ctx).Where("id = ?").Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
