package service

import (
	"context"

	"github.com/Baojiazhong/dousheng-ubuntu/cmd/user/dal/db"
	"github.com/Baojiazhong/dousheng-ubuntu/kitex_gen/userdemo"
)

type UpdateUserService struct {
	ctx context.Context
}

// NewUpdateUserService new UpdateUserService
func NewUpdateUserService(ctx context.Context) *UpdateUserService {
	return &UpdateUserService{ctx: ctx}
}

func (s *UpdateUserService) UpdateUser(req *userdemo.UpdateUserRequest) error {
	err := db.UpdateUser(s.ctx, req)
	return err
}
