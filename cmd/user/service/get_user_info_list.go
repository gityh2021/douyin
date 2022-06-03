package service

import (
	"context"
	"douyin/v1/cmd/user/dal/db"
	"douyin/v1/cmd/user/pack"
	"douyin/v1/kitex_gen/user"
)

type GetUserInfoListService struct {
	ctx context.Context
}

func NewGetUserInfoListService(ctx context.Context) *GetUserInfoListService {
	return &GetUserInfoListService{ctx: ctx}
}

func (s *GetUserInfoListService) GetUserInfoList(req *user.GetUserInfoListRequest) ([]*user.User, error) {

	// 需要传入登录用户ID

	userModels, err := db.GetUserInfoList(s.ctx, req.UserIds)
	if err != nil {
		return nil, err
	}
	IsFollowList, err := db.QueryFollowRelation(s.ctx, userModels, req.UserId)
	if err != nil {
		return nil, err
	}

	users := pack.Users(userModels, IsFollowList)

	return users, nil
}
