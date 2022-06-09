package handlers

import (
	"context"
	"douyin/v1/cmd/api/vo"
	"strconv"

	"douyin/v1/cmd/api/rpc"
	"douyin/v1/kitex_gen/user"
	"douyin/v1/pkg/constants"
	"douyin/v1/pkg/errno"

	"github.com/gin-gonic/gin"
)

//GetFollowList 获取关注列表
func GetFollowList(c *gin.Context) {
	tokenId := vo.GetUserIdFromToken(c)
	toUserIDString := c.Query("user_id")
	toUserID, err := strconv.ParseInt(toUserIDString, 10, 64)
	req := &user.MGetUserRequest{UserId: tokenId, ToUserId: toUserID, ActionType: constants.QueryFollowList}
	user, err := rpc.MGetUser(context.Background(), req)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendUserListResponse(c, errno.Success, user)
}

// GetFollowerList 获取粉丝列表
func GetFollowerList(c *gin.Context) {
	tokenId := vo.GetUserIdFromToken(c)
	toUserIDString := c.Query("user_id")
	toUserID, err := strconv.ParseInt(toUserIDString, 10, 64)

	req := &user.MGetUserRequest{UserId: tokenId, ToUserId: toUserID, ActionType: constants.QueryFollowerList}
	user, err := rpc.MGetUser(context.Background(), req)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendUserListResponse(c, errno.Success, user)
}
