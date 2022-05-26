package db

import (
	"context"

	"github.com/Baojiazhong/dousheng-ubuntu/kitex_gen/userdemo"
	"github.com/Baojiazhong/dousheng-ubuntu/pkg/constants"
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

func DealWithFollowRelation(ctx context.Context, req *userdemo.UpdateUserRequest) error {
	if req.ActionType == constants.RelationAdd {
		// 添加关注
		return DB.WithContext(ctx).Create(&Follower{UserID: req.ToUserId, FollowerID: req.UserId}).Error
	} else if req.ActionType == constants.RelationDel {
		// 取消关注
		return DB.WithContext(ctx).Where("user_id = ? and follower_id = ?", req.ToUserId, req.UserId).Delete(&Follower{}).Error
	}
	return nil
}
