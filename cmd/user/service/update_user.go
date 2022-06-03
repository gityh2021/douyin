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

func (s *UpdateUserService) UpdateUser(req *user.UpdateUserRequest) error {
	err := db.UpdateUser(s.ctx, req)
	return err
}
