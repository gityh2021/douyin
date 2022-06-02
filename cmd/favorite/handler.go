package main

import (
	"context"
	"douyin/v1/cmd/favorite/pack"
	"douyin/v1/cmd/video/service"
	"douyin/v1/kitex_gen/favorite"
	"douyin/v1/pkg/errno"
)

// FavoriteServiceImpl implements the last service interface defined in the IDL.
type FavoriteServiceImpl struct{}

// FavoriteByUser implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteByUser(ctx context.Context, request *favorite.FavoriteActionRequest) (resp *favorite.BaseResponse, err error) {
	response := new(favorite.BaseResponse)

	if request.UserId < 0 {
		response = pack.BuildBaseResp(errno.ParamErr)
		return response, nil
	}
	err = service.NewQueryFavoriteService(ctx).FavoriteByUser(request.UserId, request.VideoId)
	if err != nil {
		response = pack.BuildBaseResp(err)
		return response, nil
	}
	response = pack.BuildBaseResp(errno.Success)
	return response, nil
}

// GetFavoriteListBYUser implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) GetFavoriteListBYUser(ctx context.Context, request *favorite.FavoriteListRequest) (resp *favorite.FavoriteListResponse, err error) {
	response := new(favorite.FavoriteListResponse)

	if request.UserId < 0 {
		response.SetBaseResp(pack.BuildBaseResp(errno.ParamErr))
		return response, nil
	}
	videos, err := service.NewQueryFavoriteService(ctx).GetFavoriteListByUser(request.UserId)
	if err != nil {
		response.SetBaseResp(pack.BuildBaseResp(err))
		return response, nil
	}
	response.SetBaseResp(pack.BuildBaseResp(errno.Success))
	response.SetVideoList(videos)
	return response, nil
}
