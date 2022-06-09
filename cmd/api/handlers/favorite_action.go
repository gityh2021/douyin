package handlers

import (
	"context"
	"douyin/v1/cmd/api/rpc"
	"douyin/v1/cmd/api/vo"
	"douyin/v1/kitex_gen/video"
	"douyin/v1/pkg/constants"
	"douyin/v1/pkg/errno"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FavoriteActionParam struct {
	Token      string `json:"token"`
	UserId     int64  `json:"user_id"`
	VideoId    int64  `json:"video_id"`
	ActionType int64  `json:"action_type"`
}

// FavoriteByUser 对某视频进行点赞操作
func FavoriteByUser(c *gin.Context) {
	tokenId := vo.GetUserIdFromToken(c)
	videoIdStr := c.Query("video_id")
	actionTypeStr := c.Query("action_type")
	if tokenId == -1 || videoIdStr == "" || actionTypeStr == "" {
		SendResponse(c, errno.ParamErr, nil)
		return
	}
	if tokenId == constants.NotLogin {
		SendResponse(c, errno.LoginErr, nil)
		return
	}
	videoId, err := strconv.ParseInt(videoIdStr, 10, 64)
	if err != nil {
		SendResponse(c, err, nil)
		return
	}
	actionType, err := strconv.ParseInt(actionTypeStr, 10, 64)
	if err != nil {
		SendResponse(c, err, nil)
		return
	}
	request := video.NewFavoriteActionRequest()
	request.UserId = tokenId
	request.VideoId = videoId
	request.ActionType = actionType
	resp, err := rpc.FavoriteByUser(context.Background(), request)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, resp)
}
