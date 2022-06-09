package service

import (
	"context"
	"douyin/v1/cmd/user/dal/db"
	"douyin/v1/cmd/user/pack"
	"douyin/v1/kitex_gen/user"
)

type MGetUserService struct {
	ctx context.Context
}

// NewMGetUserService new MGetUserService
func NewMGetUserService(ctx context.Context) *MGetUserService {
	return &MGetUserService{ctx: ctx}
}

// MGetUser 根据传入user_id查询其关注/粉丝列表，同时查询登录用户是否已关注这些用户（未登录默认未关注）。
func (s *MGetUserService) MGetUser(req *user.MGetUserRequest) ([]*user.User, error) {
	Users, err := db.MGetUsers(s.ctx, req)
	if err != nil {
		return nil, err
	}
	isFollowList, err1 := db.QueryFollowRelation(s.ctx, Users, req.UserId)
	if err1 != nil {
		return nil, err1
	}
	return pack.Users(Users, isFollowList), nil
}
