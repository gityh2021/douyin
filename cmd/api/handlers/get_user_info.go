package handlers

import (
	"context"

	"douyin/v1/cmd/api/rpc"
	"douyin/v1/kitex_gen/user"
	"douyin/v1/pkg/constants"
	"douyin/v1/pkg/errno"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func GetUserInfo(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	// fmt.Printf("claims: %v\n", claims)
	// fmt.Printf("claims[constants.IdentityKey]: %v\n", claims[constants.IdentityKey])
	userID := int64(claims[constants.IdentityKey].(float64))

	req := &user.InfoGetUserRequest{UserId: userID}
	user, err := rpc.InfoGetUser(context.Background(), req)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendUserInfoResponse(c, errno.Success, user)
}
