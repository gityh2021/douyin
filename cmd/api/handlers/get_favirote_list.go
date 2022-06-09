package handlers

import (
	"context"
	"douyin/v1/cmd/api/rpc"
	"douyin/v1/cmd/api/vo"
	"douyin/v1/pkg/errno"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FavoriteListParam struct {
	Token  string `json:"token"`
	UserId int64  `json:"user_id"`
}

//GetFavoriteList 获取喜欢的视频列表
func GetFavoriteList(c *gin.Context) {
	userIdStr := c.Query("user_id")
	if userIdStr == "" {
		SendQueryByVideoList(c, errno.ParamErr, nil)
		return
	}
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		SendQueryByVideoList(c, errno.ConvertErr(err), nil)
		return
	}
	tokenId := vo.GetUserIdFromToken(c)

	videos, err := rpc.GetFavoriteList(context.Background(), userId)
	if err != nil {
		SendQueryByVideoList(c, errno.ConvertErr(err), nil)
		return
	}
	if len(videos) == 0 {
		SendQueryByVideoList(c, errno.Success, videos)
		return
	} else {
		ids := make([]int64, len(videos))
		for i := 0; i < len(videos); i++ {
			ids[i] = videos[i].AuthorId
			videos[i].IsFavorite = true
		}
		users, err := rpc.GetUsersByIds(c, ids, tokenId)
		if err != nil {
			SendQueryByVideoList(c, err, nil)
			return
		}
		videoVos := vo.PackVideoVos(users, videos)
		SendQueryByVideoList(c, errno.Success, videoVos)
	}
}
