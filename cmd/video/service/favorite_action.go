package service

import (
	"douyin/v1/cmd/video/dal/db"
	"douyin/v1/pkg/errno"
)

func (s *QueryFavoriteService) FavoriteByUser(userId int64, videoId int64, actionType int64) error {
	switch actionType {
	case 1:
		return db.MPostFavoriteAction(s.ctx, userId, videoId)
	case 2:
		return db.MCancelFavoriteAction(s.ctx, userId, videoId)
	}
	return errno.ActionUnSupportErr
}
