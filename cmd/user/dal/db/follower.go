package db

import (
	"context"

	"douyin/v1/pkg/constants"

	"gorm.io/gorm"
)

type Follower struct {
	gorm.Model
	UserID     int64 `gorm:"not null" json:"user_id"`     // not use foreign key
	FollowerID int64 `gorm:"not null" json:"follower_id"` // not use foreign key
}

func (u *Follower) TableName() string {
	return constants.FollowerTableName
}

//QueryFollowById 根据ID查询关注列表
func QueryFollowById(ctx context.Context, userId int64) ([]*int64, error) {
	res := make([]*Follower, 0)
	if err := DB.WithContext(ctx).Where("follower_id = ?", userId).Find(&res).Error; err != nil {
		return nil, err
	}
	followIds := make([]*int64, 0)
	for _, v := range res {
		followIds = append(followIds, &v.UserID)
	}
	return followIds, nil
}

//QueryFollowerById 根据ID查询粉丝列表
func QueryFollowerById(ctx context.Context, userId int64) ([]*int64, error) {
	res := make([]*Follower, 0)
	if err := DB.WithContext(ctx).Where("user_id = ?", userId).Find(&res).Error; err != nil {
		return nil, err
	}
	followerIds := make([]*int64, 0)
	for _, v := range res {
		followerIds = append(followerIds, &v.FollowerID)
	}
	return followerIds, nil
}
