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

// NewGetUserInfoListService new GetUserInfoListService
func NewGetUserInfoListService(ctx context.Context) *GetUserInfoListService {
	return &GetUserInfoListService{ctx: ctx}
}

// GetUserInfoList 根据用户ID列表批量获取用户信息，同时查询登录用户是否已关注这些用户。用于获取视频作者列表、用户关注与粉丝列表。
func (s *GetUserInfoListService) GetUserInfoList(req *user.GetUserInfoListRequest) ([]*user.User, error) {

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
