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

// func DealWithFollowRelation(ctx context.Context, req *user.UpdateUserRequest) error {
// 	if req.UserId == constants.NotLogin {
// 		return nil
// 	}

// 	if req.ActionType == constants.RelationAdd {
// 		// 添加关注
// 		var cnt int64 = 0
// 		err := DB.WithContext(ctx).Model(&Follower{}).Where("user_id = ? and follower_id = ?", req.ToUserId, req.UserId).Count(&cnt).Error
// 		if err != nil {
// 			return err
// 		}
// 		if cnt > 0 {
// 			return nil
// 		}
// 		return DB.WithContext(ctx).Create(&Follower{UserID: req.ToUserId, FollowerID: req.UserId}).Error

// 	} else if req.ActionType == constants.RelationDel {
// 		// 取消关注
// 		return DB.WithContext(ctx).Where("user_id = ? and follower_id = ?", req.ToUserId, req.UserId).Delete(&Follower{}).Error
// 	}
// 	return nil
// }
