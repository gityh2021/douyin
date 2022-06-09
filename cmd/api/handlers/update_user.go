package handlers

import (
	"context"
	"douyin/v1/cmd/api/rpc"
	"douyin/v1/kitex_gen/user"
	"douyin/v1/pkg/constants"
	"douyin/v1/pkg/errno"
	"douyin/v1/pkg/myjwt"
	"strconv"

	"github.com/gin-gonic/gin"
)

//FollowAction 对某用户进行关注操作
func FollowAction(c *gin.Context) {
	claims := myjwt.ExtractClaims(c)
	userID := int64(claims[constants.IdentityKey].(float64))
	toUserID, err := strconv.ParseInt(c.Query("to_user_id"), 10, 64)

	if userID == toUserID {
		SendResponse(c, errno.FollowSelfErr, nil)
		return
	}

	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	actionType, err := strconv.Atoi(c.Query("action_type"))
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	req := &user.UpdateUserRequest{UserId: userID, ToUserId: int64(toUserID), ActionType: int32(actionType)}
	err = rpc.UpdateUser(context.Background(), req)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, nil)
}
