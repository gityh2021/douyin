package handlers

import (
	"douyin/v1/pkg/errno"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int64       `json:"code"`
	Message string      `json:"status_message"`
	Data    interface{} `json:"video_list"`
}

// SendResponse pack response
func SendResponse(c *gin.Context, err error, data interface{}) {
	Err := errno.ConvertErr(err)
	c.JSON(http.StatusOK, Response{
		Code:    Err.ErrCode,
		Message: Err.ErrMsg,
		Data:    data,
	})
}
