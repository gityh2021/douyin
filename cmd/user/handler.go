package main

import (
	"context"

	"github.com/Baojiazhong/dousheng-ubuntu/cmd/user/pack"
	"github.com/Baojiazhong/dousheng-ubuntu/cmd/user/service"
	"github.com/Baojiazhong/dousheng-ubuntu/kitex_gen/userdemo"
	"github.com/Baojiazhong/dousheng-ubuntu/pkg/errno"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// CreateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CreateUser(ctx context.Context, req *userdemo.CreateUserRequest) (resp *userdemo.CreateUserResponse, err error) {
	// TODO: Your code here...
	resp = new(userdemo.CreateUserResponse)

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

// CheckUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CheckUser(ctx context.Context, req *userdemo.CheckUserRequest) (resp *userdemo.CheckUserResponse, err error) {
	// TODO: Your code here...
	resp = new(userdemo.CheckUserResponse)

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

// InfoGetUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) InfoGetUser(ctx context.Context, req *userdemo.InfoGetUserRequest) (resp *userdemo.InfoGetUserResponse, err error) {
	// TODO: Your code here...
	resp = new(userdemo.InfoGetUserResponse)

	if req.UserId < 1 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	user, err := service.NewInfoGetUserService(ctx).InfoGetUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	// resp.User.UserId = user.ID
	// resp.User.Username = user.UserName
	// resp.User.FollowCount = user.FollowCount
	// resp.User.FollowerCount = user.FollowerCount
	resp.User = pack.User(user, false) // should not do here
	resp.BaseResp = pack.BuildBaseResp(errno.Success)

	return
}

// MGetUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) MGetUser(ctx context.Context, req *userdemo.MGetUserRequest) (resp *userdemo.MGetUserResponse, err error) {
	// TODO: Your code here...
	// user info, follow and follower
	resp = new(userdemo.MGetUserResponse)

	if req.UserId < 1 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	users, err := service.NewMGetUserService(ctx).MGetUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.Users = users
	return resp, nil
}

// UpdateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdateUser(ctx context.Context, req *userdemo.UpdateUserRequest) (resp *userdemo.UpdateUserResponse, err error) {
	// TODO: Your code here...
	resp = new(userdemo.UpdateUserResponse)

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
