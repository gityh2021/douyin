package handlers

import (
	"context"
	"douyin/v1/cmd/api/oss"
	"douyin/v1/cmd/api/rpc"
	"douyin/v1/cmd/api/vo"
	"douyin/v1/kitex_gen/video"
	"douyin/v1/pkg/constants"
	"douyin/v1/pkg/errno"
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
	videoFilename := "video/" + strconv.FormatInt(userID, 10) + strconv.FormatInt(time.Now().Unix(), 10) + filepath.Base(data.Filename)
	coverFileName := "cover/" + strconv.FormatInt(userID, 10) + strconv.FormatInt(time.Now().Unix(), 10) + ".jpeg"
	saveFile := filepath.Join("./cmd/api/static/", videoFilename)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		SendCreateVideoResponse(c, err)
		return
	}
	v := oss.VideoOss{
		SourceFile: saveFile,
		TargetFile: videoFilename,
	}
	oss.VideoList.Lock()
	oss.VideoList.M = append(oss.VideoList.M, v)
	oss.VideoList.Unlock()
	i := oss.ImgOss{
		VideoFile:  saveFile,
		SourceFile: filepath.Join("./cmd/api/static/", coverFileName),
		TargetFile: coverFileName,
	}
	oss.ImgList.Lock()
	oss.ImgList.M = append(oss.ImgList.M, i)
	oss.ImgList.Unlock()
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
		return
	}
	SendCreateVideoResponse(c, errno.Success)
}
