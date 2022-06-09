package handlers

import (
	"context"

	"douyin/v1/cmd/api/rpc"
	"douyin/v1/kitex_gen/user"
	"douyin/v1/pkg/constants"
	"douyin/v1/pkg/errno"

	"douyin/v1/pkg/myjwt"

	"github.com/gin-gonic/gin"
)

// GetUserInfo 获取用户信息
func GetUserInfo(c *gin.Context) {
	claims := myjwt.ExtractClaims(c)
	userID := int64(claims[constants.IdentityKey].(float64))

	req := &user.InfoGetUserRequest{UserId: userID}
	user, err := rpc.InfoGetUser(context.Background(), req)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendUserInfoResponse(c, errno.Success, user)
}
