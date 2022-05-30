package handlers

import (
	"context"

	"github.com/Baojiazhong/dousheng-ubuntu/cmd/api/rpc"
	"github.com/Baojiazhong/dousheng-ubuntu/kitex_gen/userdemo"
	"github.com/Baojiazhong/dousheng-ubuntu/pkg/errno"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var registerVar UserParam
	registerVar.UserName = c.Query("username")
	registerVar.PassWord = c.Query("password")

	if len(registerVar.UserName) == 0 || len(registerVar.PassWord) == 0 {
		SendResponse(c, errno.ParamErr, nil)
		return
	}
	// kong context
	err := rpc.CreateUser(context.Background(), &userdemo.CreateUserRequest{
		Username: registerVar.UserName,
		Password: registerVar.PassWord,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	// auto login
	// authMiddleware.LoginHandler(c)
}
