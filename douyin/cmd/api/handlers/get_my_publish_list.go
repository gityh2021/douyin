package handlers

import (
	"context"
	"douyin/v1/cmd/api/rpc"
	"douyin/v1/pkg/errno"
	"github.com/gin-gonic/gin"
	"strconv"
)

type PublishParam struct {
	Token  string `json:"token"`
	UserId int64  `json:"user_id"`
}

func GetMyPublishVideoList(c *gin.Context) {
	token := c.Query("user_id")
	userIdStr := c.Query("user_id")
	if token == "" || userIdStr == "" {
		SendResponse(c, errno.ParamErr, nil)
		return
	}
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	resp, err := rpc.GetPublishVideoList(context.Background(), userId)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, resp)
}
