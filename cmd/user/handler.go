package main

import (
	"context"

	"douyin/v1/cmd/user/pack"
	"douyin/v1/cmd/user/service"
	"douyin/v1/kitex_gen/user"
	"douyin/v1/pkg/errno"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// CreateUser 检查用户名与密码并创建一个新用户。
func (s *UserServiceImpl) CreateUser(ctx context.Context, req *user.CreateUserRequest) (resp *user.CreateUserResponse, err error) {
	resp = new(user.CreateUserResponse)

	if len(req.Username) == 0 || len(req.Password) == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	err = service.NewCreateUserService(ctx).CreateUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// CheckUser 检查用户名与密码并登录。
func (s *UserServiceImpl) CheckUser(ctx context.Context, req *user.CheckUserRequest) (resp *user.CheckUserResponse, err error) {
	resp = new(user.CheckUserResponse)

	if len(req.Username) == 0 || len(req.Password) == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	uid, err := service.NewCheckUserService(ctx).CheckUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.UserId = uid
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// InfoGetUser 根据用户ID查询登录用户相关信息。
func (s *UserServiceImpl) InfoGetUser(ctx context.Context, req *user.InfoGetUserRequest) (resp *user.InfoGetUserResponse, err error) {
	resp = new(user.InfoGetUserResponse)

	if req.UserId < 1 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	userinfo, err := service.NewInfoGetUserService(ctx).InfoGetUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.User = pack.User(userinfo, false)
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return
}

// MGetUser 根据传入user_id查询其关注/粉丝列表，同时查询登录用户是否已关注这些用户（未登录默认未关注）。
func (s *UserServiceImpl) MGetUser(ctx context.Context, req *user.MGetUserRequest) (resp *user.MGetUserResponse, err error) {
	resp = new(user.MGetUserResponse)

	users, err := service.NewMGetUserService(ctx).MGetUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.Users = users
	return resp, nil
}

// UpdateUser 根据传入user_id与token执行关注/取关操作。
func (s *UserServiceImpl) UpdateUser(ctx context.Context, req *user.UpdateUserRequest) (resp *user.UpdateUserResponse, err error) {
	resp = new(user.UpdateUserResponse)

	if req.UserId < 1 || req.ToUserId < 1 || req.ActionType < 1 || req.ActionType > 2 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	err = service.NewUpdateUserService(ctx).UpdateUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// GetUserInfoList 根据用户ID列表批量获取用户信息，同时查询登录用户是否已关注这些用户。用于获取视频作者列表、用户关注与粉丝列表。
func (s *UserServiceImpl) GetUserInfoList(ctx context.Context, req *user.GetUserInfoListRequest) (resp *user.GetUserInfoListResponse, err error) {
	resp = new(user.GetUserInfoListResponse)

	if len(req.UserIds) == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	users, err := service.NewGetUserInfoListService(ctx).GetUserInfoList(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.Users = users
	return resp, nil
}
