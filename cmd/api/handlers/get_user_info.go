package handlers

import (
	"context"

	"github.com/Baojiazhong/dousheng-ubuntu/cmd/api/rpc"
	"github.com/Baojiazhong/dousheng-ubuntu/kitex_gen/userdemo"
	"github.com/Baojiazhong/dousheng-ubuntu/pkg/constants"
	"github.com/Baojiazhong/dousheng-ubuntu/pkg/errno"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func GetUserInfo(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	// fmt.Printf("claims: %v\n", claims)
	// fmt.Printf("claims[constants.IdentityKey]: %v\n", claims[constants.IdentityKey])
	userID := int64(claims[constants.IdentityKey].(float64))

	req := &userdemo.InfoGetUserRequest{UserId: userID}
	user, err := rpc.InfoGetUser(context.Background(), req)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendUserInfoResponse(c, errno.Success, user.UserId, user.Username, user.FollowCount, user.FollowerCount, false)
}
