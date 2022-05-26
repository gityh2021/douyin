package handlers

import (
	"net/http"
	"time"

	"github.com/Baojiazhong/dousheng-ubuntu/kitex_gen/userdemo"
	"github.com/Baojiazhong/dousheng-ubuntu/pkg/errno"
	"github.com/gin-gonic/gin"
)

type BaseResponse struct {
	Code    int32  `json:"status_code"`
	Message string `json:"status_msg"`
}

type Response struct {
	Code    int32       `json:"status_code"`
	Message string      `json:"status_msg"`
	Data    interface{} `json:"data"`
}

type LoginResponse struct {
	Code     int32  `json:"status_code"`
	Message  string `json:"status_msg"`
	UserId   int64  `json:"user_id"`
	UserName string `json:"name"`
	Token    string `json:"token"`
	Expire   string `json:"expire"`
}

type UserInfoResponse struct {
	Code          int32  `json:"status_code"`
	Message       string `json:"status_msg"`
	UserId        int64  `json:"id"`
	UserName      string `json:"name"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}

type UserListResponse struct {
	Code     int32       `json:"status_code"`
	Message  string      `json:"status_msg"`
	UserList interface{} `json:"user_list"`
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

// SendLoginResponse pack response
func SendLoginResponse(c *gin.Context, err error, userId int64, userName string, token string, expire time.Time) {
	Err := errno.ConvertErr(err)
	c.JSON(http.StatusOK, LoginResponse{
		Code:     Err.ErrCode,
		Message:  Err.ErrMsg,
		UserId:   userId,
		UserName: userName,
		Token:    token,
		Expire:   expire.Format(time.RFC3339),
	})
}

// SendUserInfoResponse pack response
func SendUserInfoResponse(c *gin.Context, err error, userId int64, userName string, followCount int64, followerCount int64, isFollow bool) {
	Err := errno.ConvertErr(err)
	c.JSON(http.StatusOK, UserInfoResponse{
		Code:          Err.ErrCode,
		Message:       Err.ErrMsg,
		UserId:        userId,
		UserName:      userName,
		FollowCount:   followCount,
		FollowerCount: followerCount,
		IsFollow:      isFollow,
	})
}

// SendUserListResponse
func SendUserListResponse(c *gin.Context, err error, userList []*userdemo.User) {
	Err := errno.ConvertErr(err)
	users := map[string]interface{}{"user_list": userList}
	c.JSON(http.StatusOK, UserListResponse{
		Code:     Err.ErrCode,
		Message:  Err.ErrMsg,
		UserList: users["user_list"],
	})
}

type UserParam struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
}
