package db

import (
	"context"
	videodb "douyin/v1/cmd/video/dal/db"
	"fmt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID            int64           `gorm:"primarykey" json:"id"`
	Name          string          `gorm:"not null" json:"name"`
	FollowCount   int64           `gorm:"default:0" json:"follow_count"`
	FollowerCount int64           `gorm:"default:0" json:"follower_count"`
	IsFollow      bool            `json:"is_follow"`
	FavoriteVideo []videodb.Video `gorm:"many2many:users_videos;" json:"favorite_video"`
}

func MGetFavoriteList(ctx context.Context, userId int64) ([]*videodb.Video, error) {
	res := make([]*videodb.Video, 0)
	if err := DB.WithContext(ctx).Model(User{}).Where("userId = ?", userId).Association("videos").Find(&res); err != nil {
		fmt.Println(err)
		return nil, err
	}
	return res, nil
}

type FavoriteVideoParam struct {
	UserId  int64
	VideoId int64
}

func MPostFavoriteAction(ctx context.Context, userId int64, videoId int64) error {
	favoriteInput := FavoriteVideoParam{
		UserId:  userId,
		VideoId: videoId,
	}
	if err := DB.WithContext(ctx).Model(User{}).Association("videos").Append(&favoriteInput); err != nil {
		return err
	}
	return nil
}
