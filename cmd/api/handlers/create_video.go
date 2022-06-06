package handlers

import (
	"context"
	"douyin/v1/cmd/api/rpc"
	"douyin/v1/cmd/api/vo"
	"douyin/v1/kitex_gen/video"
	"douyin/v1/pkg/constants"
	"douyin/v1/pkg/errno"
	"fmt"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func PublishVideo(c *gin.Context) {
	//claims := myjwt.ExtractClaims(c)
	//userID := int64(claims[constants.IdentityKey].(float64))
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
	filename := filepath.Base(data.Filename)
	finalName := fmt.Sprintf(filename)
	// 存储视频文件
	saveFile := filepath.Join("./cmd/api/static/videos/", finalName)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		SendCreateVideoResponse(c, err)
		return
	}
	newVideo := video.Video{
		AuthorId:      userID,
		PlayUrl:       constants.PlayURL + filename,
		CoverUrl:      constants.CoverURL + "1.png",
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
