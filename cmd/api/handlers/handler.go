package handlers

import (
	"douyin/v1/pkg/errno"
	"github.com/gin-gonic/gin"
	"net/http"
)

type QueryByUserIdResponse struct {
	Code    int64       `json:"status_code"`
	Message string      `json:"status_msg"`
	Data    interface{} `json:"video_list"`
}

func SendQueryByUserIdResponse(c *gin.Context, err error, data interface{}) {
	Err := errno.ConvertErr(err)
	c.JSON(http.StatusOK, QueryByUserIdResponse{
		Code:    Err.ErrCode,
		Message: Err.ErrMsg,
		Data:    data,
	})
}

type QueryByLastTimeResponse struct {
	Code     int64       `json:"status_code"`
	Message  string      `json:"status_msg"`
	NextTime int64       `json:"next_time"`
	Data     interface{} `json:"video_list"`
}

func SendQueryByLastTimeResponse(c *gin.Context, err error, data interface{}, nextTime int64) {
	Err := errno.ConvertErr(err)
	c.JSON(http.StatusOK, QueryByLastTimeResponse{
		Code:     Err.ErrCode,
		Message:  Err.ErrMsg,
		NextTime: nextTime,
		Data:     data,
	})
}

type CreateVideoResponse struct {
	Code    int64  `json:"status_code"`
	Message string `json:"status_msg"`
}

func SendCreateVideoResponse(c *gin.Context, err error) {
	Err := errno.ConvertErr(err)
	c.JSON(http.StatusOK, CreateVideoResponse{
		Code:    Err.ErrCode,
		Message: Err.ErrMsg,
	})
}

type VideoCommentParam struct {
	CommentId     int64  `json:"comment_id"`
	CommentUserId int64  `json:"comment_user_id"`
	Content       string `json:"content"`
	CreateDate    string `json:"create_date"`
}
