package handlers

import (
	"context"
	"douyin/v1/cmd/api/rpc"
	"douyin/v1/cmd/api/vo"
	"douyin/v1/pkg/errno"
	"github.com/gin-gonic/gin"
	"strconv"
)

type FavoriteListParam struct {
	Token  string `json:"token"`
	UserId int64  `json:"user_id"`
}

func GetFavoriteLIst(c *gin.Context) {
	token := c.Query("token")
	userIdStr := c.Query("user_id")
	if token == "" || userIdStr == "" {
		SendResponse(c, errno.ParamErr, nil)
		return
	}
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	tokenId := vo.GetUserIdFromToken(c)
	if tokenId == -1 {
		SendResponse(c, errno.LoginErr, nil)
		return
	}
	if tokenId != userId {
		SendResponse(c, errno.IdNotEqualErr, nil)
		return
	}
	videos, err := rpc.GetFavoriteList(context.Background(), tokenId)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	if len(videos) == 0 {
		SendResponse(c, errno.Success, videos)
		return
	} else {
		ids := make([]int64, len(videos))
		for i := 0; i < len(videos); i++ {
			ids[i] = videos[i].AuthorId
			videos[i].IsFavorite = true
		}
		users, err := rpc.GetUsersByIds(c, ids, tokenId)
		if err != nil {
			SendResponse(c, err, nil)
			return
		}
		videoVos := vo.PackVideoVos(users, videos)
		SendResponse(c, errno.Success, videoVos)
	}
}
