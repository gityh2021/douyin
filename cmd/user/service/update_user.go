package service

import (
	"context"

	"douyin/v1/cmd/user/dal/db"
	"douyin/v1/kitex_gen/user"
)

type UpdateUserService struct {
	ctx context.Context
}

// NewUpdateUserService new UpdateUserService
func NewUpdateUserService(ctx context.Context) *UpdateUserService {
	return &UpdateUserService{ctx: ctx}
}

// UpdateUser 根据传入user_id与token执行关注/取关操作。
func (s *UpdateUserService) UpdateUser(req *user.UpdateUserRequest) error {
	err := db.UpdateUser(s.ctx, req)
	return err
}
