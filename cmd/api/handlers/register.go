package handlers

import (
	"context"

	"douyin/v1/cmd/api/rpc"
	"douyin/v1/kitex_gen/user"
	"douyin/v1/pkg/errno"

	"github.com/gin-gonic/gin"
)

// Register 登陆路径调用的函数
func Register(c *gin.Context) {
	var registerVar UserParam
	registerVar.UserName = c.Query("username")
	registerVar.PassWord = c.Query("password")

	if len(registerVar.UserName) == 0 || len(registerVar.PassWord) == 0 {
		SendResponse(c, errno.ParamErr, nil)
		return
	}

	err := rpc.CreateUser(context.Background(), &user.CreateUserRequest{
		Username: registerVar.UserName,
		Password: registerVar.PassWord,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
}
