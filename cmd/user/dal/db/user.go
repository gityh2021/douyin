package db

import (
	"douyin/v1/kitex_gen/user"
	"douyin/v1/pkg/constants"
	"douyin/v1/pkg/errno"

	"gorm.io/gorm"

	"context"
)

type User struct {
	gorm.Model
	ID            int64  `gorm:"primarykey" json:"user_id"`
	UserName      string `gorm:"not null" json:"user_name"`       // not null and repeat
	Password      string `gorm:"not null" json:"password"`        // md5加密后的密码
	FollowCount   int64  `gorm:"default:0" json:"follow_count"`   // 关注数
	FollowerCount int64  `gorm:"default:0" json:"follower_count"` // 粉丝数
}

func (u *User) TableName() string {
	return constants.UserTableName
}

// CreateUser create user info
func CreateUser(ctx context.Context, users []*User) error {
	return DB.WithContext(ctx).Create(users).Error
}

// // MGetUsers multiple get list of user info
// func MGetUsers(ctx context.Context, userIDs []int64) ([]*User, error) {
// 	res := make([]*User, 0)
// 	if len(userIDs) == 0 {
// 		return res, nil
// 	}

// 	if err := DB.WithContext(ctx).Where("id in ?", userIDs).Find(&res).Error; err != nil {
// 		return nil, err
// 	}
// 	return res, nil
// }

// QueryUser query list of user info
func QueryUser(ctx context.Context, userName string) ([]*User, error) {
	res := make([]*User, 0)
	if err := DB.WithContext(ctx).Where("user_name = ?", userName).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// QueryUser query list of user info
func QueryUserById(ctx context.Context, userId int64) (User, error) {
	res := User{}
	if err := DB.WithContext(ctx).Where("id = ?", userId).First(&res).Error; err != nil {
		return User{}, err
	}
	return res, nil
}

// MGetUsers multiple get list of user info
func MGetUsers(ctx context.Context, req *user.MGetUserRequest) ([]*User, error) {
	res := make([]*User, 0)
	if req.UserId < 1 || req.ActionType < constants.QueryUserInfo || req.ActionType > constants.QueryFollowerList {
		return nil, errno.ParamErr
	}
	if req.ActionType == constants.QueryUserInfo {
		// query user info
		if err := DB.WithContext(ctx).Where("id = ?", req.UserId).Find(&res).Error; err != nil {
			return nil, err
		}
		return res, nil
	} else if req.ActionType == constants.QueryFollowList {
		// query follow list
		// step 1: query table follower
		followIds, err := QueryFollowById(ctx, req.UserId)
		if err != nil {
			return nil, err
		}
		// step 2: query table user
		if err = DB.WithContext(ctx).Where("id in ?", followIds).Find(&res).Error; err != nil {
			return nil, err
		}
		return res, nil
	} else {
		// if req.ActionType == constants.QueryFollowerList
		// query follower list
		// step 1: query table follower
		followerIds, err := QueryFollowerById(ctx, req.UserId)
		if err != nil {
			return nil, err
		}
		// step 2: query table user
		if err = DB.WithContext(ctx).Where("id in ?", followerIds).Find(&res).Error; err != nil {
			return nil, err
		}
		return res, nil
	}
}

//UpdateUser需要使用事务
func UpdateUser(ctx context.Context, req *user.UpdateUserRequest) error {
	// step 1: query table follower
	err := DealWithFollowRelation(ctx, req)
	if err != nil {
		return err
	}
	// step 2: update table user
	var user1 User
	var user2 User
	//这里是关注
	if req.ActionType == constants.RelationAdd {
		var cnt int64 = 0
		err := DB.WithContext(ctx).Model(&Follower{}).Where("user_id = ? and follower_id = ?", req.ToUserId, req.UserId).Count(&cnt).Error
		if err != nil {
			return err
		}
		if cnt > 0 {
			return nil
		}

		if err = DB.WithContext(ctx).Where("id = ?", req.UserId).First(&user1).Error; err != nil {
			return err
		}

		if err = DB.WithContext(ctx).Where("id = ?", req.ToUserId).First(&user2).Error; err != nil {
			return err
		}

		user1.FollowCount += 1
		user2.FollowerCount += 1
		DB.Transaction(func(tx *gorm.DB) error {
			// 在事务中执行一些 db 操作（从这里开始，应该使用 'tx' 而不是 'db'）
			if err := tx.WithContext(ctx).Model(&user1).Select("follow_count").Updates(user1).Error; err != nil {
				// 返回任何错误都会回滚事务
				return err
			}

			if err := tx.WithContext(ctx).Model(&user2).Select("follower_count").Updates(user2).Error; err != nil {
				return err
			}
			// 返回 nil 提交事务
			return nil
		}) //下面是取关
	} else if req.ActionType == constants.RelationDel {
		if err = DB.WithContext(ctx).Where("id = ?", req.UserId).First(&user1).Error; err != nil {
			return err
		}

		if err = DB.WithContext(ctx).Where("id = ?", req.ToUserId).First(&user2).Error; err != nil {
			return err
		}

		user1.FollowCount -= 1
		user2.FollowerCount -= 1
		DB.Transaction(func(tx *gorm.DB) error {
			// 在事务中执行一些 db 操作（从这里开始，应该使用 'tx' 而不是 'db'）
			if err := tx.WithContext(ctx).Model(&user1).Select("follow_count").Updates(user1).Error; err != nil {
				// 返回任何错误都会回滚事务
				return err
			}

			if err := tx.WithContext(ctx).Model(&user2).Select("follower_count").Updates(user2).Error; err != nil {
				return err
			}
			// 返回 nil 提交事务
			return nil
		}) //下面是取关

	}
	return nil
}

func QueryFollowRelation(ctx context.Context, users []*User, userId int64) ([]bool, error) {
	isFollowList := make([]bool, len(users))
	if userId == constants.NotLogin {
		for i := 0; i < len(users); i++ {
			isFollowList[i] = false
		}

	} else {
		for i, user := range users {
			var temp int64 = 0
			DB.WithContext(ctx).Model(&Follower{}).Where("user_id = ? and follower_id = ?", user.ID, userId).Count(&temp)
			if temp > 0 {
				isFollowList[i] = true
			} else {
				isFollowList[i] = false
			}
		}
	}
	return isFollowList, nil
}

func GetUserInfoList(ctx context.Context, userIDs []int64) ([]*User, error) {
	var res []*User
	if len(userIDs) == 0 {
		return res, nil
	}

	if err := DB.WithContext(ctx).Where("id in ?", userIDs).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
