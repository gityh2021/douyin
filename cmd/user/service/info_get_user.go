package service

import (
	"context"

	"github.com/Baojiazhong/dousheng-ubuntu/cmd/user/dal/db"
	"github.com/Baojiazhong/dousheng-ubuntu/kitex_gen/userdemo"
	"github.com/Baojiazhong/dousheng-ubuntu/pkg/errno"
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
func (s *InfoGetUserService) InfoGetUser(req *userdemo.InfoGetUserRequest) (*db.User, error) {
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
