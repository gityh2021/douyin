package main

import (
	"context"

	"github.com/Baojiazhong/dousheng-ubuntu/kitex_gen/userdemo"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// CreateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CreateUser(ctx context.Context, req *userdemo.CreateUserRequest) (resp *userdemo.CreateUserResponse, err error) {
	// TODO: Your code here...
	return
}

// CheckUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CheckUser(ctx context.Context, req *userdemo.CheckUserRequest) (resp *userdemo.CheckUserResponse, err error) {
	// TODO: Your code here...
	return
}

// InfoGetUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) InfoGetUser(ctx context.Context, req *userdemo.InfoGetUserRequest) (resp *userdemo.InfoGetUserResponse, err error) {
	// TODO: Your code here...
	return
}

// MGetUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) MGetUser(ctx context.Context, req *userdemo.MGetUserRequest) (resp *userdemo.MGetUserResponse, err error) {
	// TODO: Your code here...
	return
}

// UpdateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdateUser(ctx context.Context, req *userdemo.UpdateUserRequest) (resp *userdemo.UpdateUserResponse, err error) {
	// TODO: Your code here...
	return
}
