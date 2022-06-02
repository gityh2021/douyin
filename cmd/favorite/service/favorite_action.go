package service

import (
	"douyin/v1/cmd/favorite/dal/db"
)

func (s *QueryFavoriteService) FavoriteByUser(userId int64, videoId int64) error {
	err := db.MPostFavoriteAction(s.ctx, userId, videoId)
	if err != nil {
		return err
	}
	return nil
}
