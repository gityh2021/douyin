package handlers

import (
	"context"
	"douyin/v1/cmd/api/rpc"
	"douyin/v1/pkg/errno"
	"github.com/gin-gonic/gin"
	"strconv"
)

type VideoCommentsParam struct {
	Token  string `json:"token"`
	VideoId int64  `json:"video_id"`
}

func GetVideoCommentsList(c *gin.Context) {
	token := c.Query("video_id")
	videoIdStr := c.Query("video_id")
	if token == "" || videoIdStr == "" {
		SendResponse(c, errno.ParamErr, nil)
		return
	}
	videoId, err := strconv.ParseInt(videoIdStr, 10, 64)
	resp, err := rpc.GetVideoCommentsList(context.Background(), videoId)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, resp)
}
