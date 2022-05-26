package service

import (
	"context"
	"fmt"

	"github.com/Baojiazhong/dousheng-ubuntu/cmd/user/dal/db"
	"github.com/Baojiazhong/dousheng-ubuntu/cmd/user/pack"
	"github.com/Baojiazhong/dousheng-ubuntu/kitex_gen/userdemo"
	"github.com/Baojiazhong/dousheng-ubuntu/pkg/constants"
)

type MGetUserService struct {
	ctx context.Context
}

// NewMGetUserService new MGetUserService
func NewMGetUserService(ctx context.Context) *MGetUserService {
	return &MGetUserService{ctx: ctx}
}

// MGetUser multiple get list of user info
func (s *MGetUserService) MGetUser(req *userdemo.MGetUserRequest) ([]*userdemo.User, error) {
	Users, err := db.MGetUsers(s.ctx, req)
	if err != nil {
		return nil, err
	}
	if req.ActionType == constants.QueryFollowList {
		// 查询到的用户列表中的is_follow字段一定为true
		isFollowList := make([]int64, len(Users))
		for i := range Users {
			isFollowList[i] = 1
		}
		return pack.Users(Users, isFollowList), nil
	} else if req.ActionType == constants.QueryFollowerList {
		// 还得确认一下是否关注了该用户
		isFollowLIst, err1 := db.QueryFollowRelation(s.ctx, Users, req.UserId)
		if err1 != nil {
			return nil, err1
		}
		for i := range isFollowLIst {
			fmt.Println("isFollowLIst[i]: ", isFollowLIst[i])
		}
		return pack.Users(Users, isFollowLIst), nil
		// return pack.Users(Users, false), nil
	} else {
		isFollowList := make([]int64, len(Users))
		for i := range Users {
			isFollowList[i] = 0
		}
		return pack.Users(Users, isFollowList), nil // default
	}
}
