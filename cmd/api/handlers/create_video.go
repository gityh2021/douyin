package handlers

import (
	"context"
	"douyin/v1/cmd/api/oss"
	"douyin/v1/cmd/api/rpc"
	"douyin/v1/cmd/api/vo"
	"douyin/v1/kitex_gen/video"
	"douyin/v1/pkg/constants"
	"douyin/v1/pkg/errno"
	"math/rand"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func PublishVideo(c *gin.Context) {
	userID := vo.GetUserIdFromToken(c)
	titleStr := c.PostForm("title")
	//如果视频标题为空 err
	if titleStr == "" {
		SendCreateVideoResponse(c, errno.ParamErr)
		return
	}

	//获取视频流
	data, err := c.FormFile("data")
	if err != nil {
		SendCreateVideoResponse(c, err)
		return
	}
	if err != nil {
		SendCreateVideoResponse(c, err)
		return
	}
	videoFilename := "video/" + strconv.FormatInt(userID, 10) + strconv.FormatInt(time.Now().Unix(), 10) + filepath.Base(data.Filename)
	if err := oss.PutFeed(data, videoFilename); err != nil {
		SendCreateVideoResponse(c, err)
		return
	}
	var coverFileName string
	if err := oss.ExampleReadFrameAsJpeg(constants.OSSFetchURL+videoFilename, 5); err != nil {
		// 随机选择一张图片作为封面
		rand.Seed(time.Now().Unix())
		coverFileName = "cover/" + strconv.Itoa(rand.Intn(20)) + ".jpeg"
	} else {
		coverFileName = "cover/" + strconv.FormatInt(userID, 10) + strconv.FormatInt(time.Now().Unix(), 10) + "out.jpeg"
	}
	if err := oss.PutImage("./out.jpeg", coverFileName); err != nil {
		SendCreateVideoResponse(c, err)
		return
	}
	newVideo := video.Video{
		AuthorId:      userID,
		PlayUrl:       constants.OSSFetchURL + videoFilename,
		CoverUrl:      constants.OSSFetchURL + coverFileName,
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
		Title:         titleStr,
	}
	if err := rpc.CreateVideo(context.Background(), &newVideo); err != nil {
		SendCreateVideoResponse(c, err)
	}
	SendCreateVideoResponse(c, errno.Success)
}
