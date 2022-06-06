package main

import (
	"context"
	"douyin/v1/kitex_gen/user"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// CreateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CreateUser(ctx context.Context, req *user.CreateUserRequest) (resp *user.CreateUserResponse, err error) {
	// TODO: Your code here...
	return
}

// CheckUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CheckUser(ctx context.Context, req *user.CheckUserRequest) (resp *user.CheckUserResponse, err error) {
	// TODO: Your code here...
	return
}

// InfoGetUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) InfoGetUser(ctx context.Context, req *user.InfoGetUserRequest) (resp *user.InfoGetUserResponse, err error) {
	// TODO: Your code here...
	return
}

// MGetUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) MGetUser(ctx context.Context, req *user.MGetUserRequest) (resp *user.MGetUserResponse, err error) {
	// TODO: Your code here...
	return
}

// UpdateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdateUser(ctx context.Context, req *user.UpdateUserRequest) (resp *user.UpdateUserResponse, err error) {
	// TODO: Your code here...
	return
}

// GetUserInfoList implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUserInfoList(ctx context.Context, req *user.GetUserInfoListRequest) (resp *user.GetUserInfoListResponse, err error) {
	// TODO: Your code here...
	return
}
