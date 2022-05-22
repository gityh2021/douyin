package main

import (
	"context"
	"douyin/v1/cmd/video/pack"
	"douyin/v1/cmd/video/service"
	"douyin/v1/kitex_gen/video"
	"douyin/v1/pkg/errno"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// GetPublishListByUser implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetPublishListByUser(ctx context.Context, userId int64) (resp *video.PublishListResponse, err error) {
	response := new(video.PublishListResponse)
	if userId < 0 {
		response.SetBaseResp(pack.BuildBaseResp(errno.ParamErr))
		return response, nil
	}
	videos, err := service.NewQueryVideoService(ctx).GetPublishList(userId)
	if err != nil {
		response.SetBaseResp(pack.BuildBaseResp(err))
		return response, nil
	}
	response.SetBaseResp(pack.BuildBaseResp(errno.Success))
	response.SetVideoList(videos)
	return response, nil
}
