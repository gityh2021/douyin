package handlers

import (
	"context"
	"douyin/pkg/errno"
	"douyin/cmd/api/rpc"
	"github.com/gin-gonic/gin"
)

type PublishParam struct {
	token 	string  `json:"token"`
	userId  int64   `json:"user_id"`
}
func getMyPublishVideoList(c *gin.Context)  {
	var publishParam PublishParam
	if err := c.ShouldBind(&publishParam); err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}

	if len(publishParam.token) == 0 || publishParam.userId < 0 {
		SendResponse(c, errno.ParamErr, nil)
		return
	}
	resp, err := rpc.GetPublishVideoList(context.Background(), publishParam.userId)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, resp)
}
