package handlers

import (
	"context"
	"douyin/v1/cmd/api/rpc"
	"douyin/v1/kitex_gen/video"
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
	request := video.NewFavoriteListRequest()
	request.UserId = userId
	request.Token = token
	resp, err := rpc.GetFavoriteList(context.Background(), request)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, resp)
}
