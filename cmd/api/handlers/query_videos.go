package handlers

import (
	"context"
	"douyin/v1/cmd/api/rpc"
	"douyin/v1/cmd/api/vo"
	"douyin/v1/pkg/errno"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

//GetPublishVideoList 获取发布视频列表
func GetPublishVideoList(c *gin.Context) {
	userIdFromToken := vo.GetUserIdFromToken(c)
	userIdStr := c.Query("user_id")
	if userIdStr == "" {
		SendQueryByVideoList(c, errno.ParamErr, nil)
		return
	}
	userIdFromQuery, err := strconv.ParseInt(userIdStr, 10, 64)
	videos, err := rpc.GetPublishVideoList(context.Background(), userIdFromQuery)
	if err != nil {
		SendQueryByVideoList(c, errno.ConvertErr(err), nil)
		return
	}

	// 将Author信息封装到VideoVo中
	if len(videos) == 0 {
		SendQueryByVideoList(c, errno.Success, videos)
		return
	} else {
		ids := make([]int64, len(videos))
		for i := 0; i < len(videos); i++ {
			ids[i] = videos[i].AuthorId
		}
		users, err := rpc.GetUsersByIds(c, ids, userIdFromToken)
		if err != nil {
			SendQueryByVideoList(c, err, nil)
			return
		}
		videoVos := vo.PackVideoVos(users, videos)
		SendQueryByVideoList(c, errno.Success, videoVos)
	}
}

func GetVideoFeed(c *gin.Context) {
	userIdFromToken := vo.GetUserIdFromToken(c)
	lastTimeStr := c.Query("latest_time")
	lastTime := time.Now().Unix()
	if lastTimeStr != "" {
		t, err := strconv.ParseInt(lastTimeStr, 10, 64)
		if err != nil {
			SendQueryByLastTimeResponse(c, errno.ConvertErr(err), nil, time.Now().Unix())
			return
		}
		if t != 0 {
			lastTime = t
		}
	}
	videos, nextTime, err := rpc.GetVideosFeed(context.Background(), lastTime, userIdFromToken)
	if err != nil {
		SendQueryByLastTimeResponse(c, err, nil, nextTime)
		return
	}

	// 将Author信息封装到VideoVo中
	if len(videos) == 0 {
		SendQueryByLastTimeResponse(c, errno.Success, videos, nextTime)
		return
	} else {
		ids := make([]int64, len(videos))
		for i := 0; i < len(videos); i++ {
			ids[i] = videos[i].AuthorId
		}
		users, err := rpc.GetUsersByIds(c, ids, userIdFromToken)
		if err != nil {
			SendQueryByLastTimeResponse(c, err, nil, nextTime)
			return
		}
		videoVos := vo.PackVideoVos(users, videos)
		SendQueryByLastTimeResponse(c, errno.Success, videoVos, nextTime)
	}
}
