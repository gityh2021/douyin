package service

import (
	"context"

	"douyin/v1/cmd/user/dal/db"
	"douyin/v1/kitex_gen/user"
	"douyin/v1/pkg/errno"
)

type InfoGetUserService struct {
	ctx context.Context
}

// NewInfoGetUserService new InfoGetUserService
func NewInfoGetUserService(ctx context.Context) *InfoGetUserService {
	return &InfoGetUserService{
		ctx: ctx,
	}
}

// InfoGetUser info get user
func (s *InfoGetUserService) InfoGetUser(req *user.InfoGetUserRequest) (*db.User, error) {
	userId := req.UserId
	user, err := db.QueryUserById(s.ctx, userId)
	if err != nil {
		return nil, err
	}
	if user.ID < 1 {
		return nil, errno.UserNotExistErr
	}
	return &user, nil
}
