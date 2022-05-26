package main

import (
	"context"
	"net/http"
	"time"

	"github.com/Baojiazhong/dousheng-ubuntu/cmd/api/handlers"
	"github.com/Baojiazhong/dousheng-ubuntu/cmd/api/rpc"
	"github.com/Baojiazhong/dousheng-ubuntu/kitex_gen/userdemo"
	"github.com/Baojiazhong/dousheng-ubuntu/pkg/constants"
	"github.com/Baojiazhong/dousheng-ubuntu/pkg/errno"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gin-gonic/gin"
)

func Init() {
	rpc.InitRPC()
}

var authMiddleware *jwt.GinJWTMiddleware

func main() {
	Init()
	r := gin.New()
	authMiddleware, _ = jwt.New(&jwt.GinJWTMiddleware{
		Key:        []byte(constants.SecretKey),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			// fmt.Printf("data: %v\n", data)
			if v, ok := data.(int64); ok {
				// fmt.Println("11111111111111111111111111111")
				return jwt.MapClaims{
					constants.IdentityKey: v,
				}
			}
			return jwt.MapClaims{}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVar1 handlers.UserParam
			loginVar1.UserName = c.Query("username")
			loginVar1.PassWord = c.Query("password")
			if len(loginVar1.UserName) == 0 || len(loginVar1.PassWord) == 0 {
				return "", jwt.ErrMissingLoginValues
			}

			return rpc.CheckUser(context.Background(), &userdemo.CheckUserRequest{Username: loginVar1.UserName, Password: loginVar1.PassWord})
		},
		LoginResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			var loginVar2 handlers.UserParam
			loginVar2.UserName = c.Query("username")
			loginVar2.PassWord = c.Query("password")
			user_id, res := rpc.QueryUser(context.Background(), &userdemo.CheckUserRequest{Username: loginVar2.UserName, Password: loginVar2.PassWord})
			// c.JSON(http.StatusOK, gin.H{
			// 	"status_code": res.ErrCode,
			// 	"status_msg":  res.ErrMsg,
			// 	"user_id":     user_id,
			// 	"token":       token,
			// })
			handlers.SendLoginResponse(c, res, user_id, loginVar2.UserName, token, expire)
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})

	v1 := r.Group("/douyin")
	user1 := v1.Group("/user")
	user1.POST("/login/", authMiddleware.LoginHandler)
	user1.POST("/register/", Register)
	// ----------------------------------------------
	user1.Use(authMiddleware.MiddlewareFunc())
	user1.GET("/", handlers.GetUserInfo)

	user2 := v1.Group("/relation")
	user2.Use(authMiddleware.MiddlewareFunc())
	user2.GET("/follow/list/", handlers.GetFollowList)
	user2.GET("/follower/list/", handlers.GetFollowerList)
	user2.POST("/action/", handlers.FollowAction)

	if err := http.ListenAndServe(":8080", r); err != nil {
		klog.Fatal(err)
	}
}

func Register(c *gin.Context) {
	var registerVar handlers.UserParam
	registerVar.UserName = c.Query("username")
	registerVar.PassWord = c.Query("password")

	if len(registerVar.UserName) == 0 || len(registerVar.PassWord) == 0 {
		handlers.SendResponse(c, errno.ParamErr, nil)
		return
	}
	// kong context
	err := rpc.CreateUser(context.Background(), &userdemo.CreateUserRequest{
		Username: registerVar.UserName,
		Password: registerVar.PassWord,
	})
	if err != nil {
		handlers.SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	// auto login
	authMiddleware.LoginHandler(c)
}
