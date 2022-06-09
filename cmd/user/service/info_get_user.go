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

// InfoGetUser 根据用户ID查询登录用户相关信息。
func (s *InfoGetUserService) InfoGetUser(req *user.InfoGetUserRequest) (*db.User, error) {
	userId := req.UserId
	userinfo, err := db.QueryUserById(s.ctx, userId)
	if err != nil {
		return nil, err
	}
	if userinfo.ID < 1 {
		return nil, errno.UserNotExistErr
	}
	return &userinfo, nil
}
