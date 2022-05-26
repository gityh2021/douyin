package handlers

import (
	"context"
	"strconv"

	"github.com/Baojiazhong/dousheng-ubuntu/cmd/api/rpc"
	"github.com/Baojiazhong/dousheng-ubuntu/kitex_gen/userdemo"
	"github.com/Baojiazhong/dousheng-ubuntu/pkg/constants"
	"github.com/Baojiazhong/dousheng-ubuntu/pkg/errno"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func FollowAction(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	userID := int64(claims[constants.IdentityKey].(float64))
	toUserID, err1 := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
	if err1 != nil {
		SendResponse(c, errno.ConvertErr(err1), nil)
		return
	}
	actionType, err2 := strconv.Atoi(c.Query("action_type"))
	if err2 != nil {
		SendResponse(c, errno.ConvertErr(err2), nil)
		return
	}
	req := &userdemo.UpdateUserRequest{UserId: userID, ToUserId: int64(toUserID), ActionType: int32(actionType)}
	err3 := rpc.UpdateUser(context.Background(), req)
	if err3 != nil {
		SendResponse(c, errno.ConvertErr(err3), nil)
		return
	}
	SendResponse(c, errno.Success, nil)
}
