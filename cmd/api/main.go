package main

import (
	"context"
	"net/http"
	"time"

	"douyin/v1/cmd/api/handlers"
	"douyin/v1/cmd/api/rpc"
	"douyin/v1/kitex_gen/user"
	"douyin/v1/pkg/constants"

	"douyin/v1/pkg/myjwt"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gin-gonic/gin"
)

func Init() {
	rpc.InitRPC()
}

func main() {
	Init()
	r := gin.Default()
	authMiddleware, _ := myjwt.New(&myjwt.GinJWTMiddleware{
		Key:        []byte(constants.SecretKey),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		PayloadFunc: func(data interface{}) myjwt.MapClaims {
			// fmt.Printf("data: %v\n", data)
			if v, ok := data.(int64); ok {
				return myjwt.MapClaims{
					constants.IdentityKey: v,
				}
			}
			return myjwt.MapClaims{}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVar1 handlers.UserParam
			loginVar1.UserName = c.Query("username")
			loginVar1.PassWord = c.Query("password")
			if len(loginVar1.UserName) == 0 || len(loginVar1.PassWord) == 0 {
				return "", myjwt.ErrMissingLoginValues
			}

			return rpc.CheckUser(context.Background(), &user.CheckUserRequest{Username: loginVar1.UserName, Password: loginVar1.PassWord})
		},
		LoginResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			var loginVar2 handlers.UserParam
			loginVar2.UserName = c.Query("username")
			loginVar2.PassWord = c.Query("password")
			user_id, res := rpc.QueryUser(context.Background(), &user.CheckUserRequest{Username: loginVar2.UserName, Password: loginVar2.PassWord})
			handlers.SendLoginResponse(c, res, user_id, loginVar2.UserName, token, expire)
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt, postform: token",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
		FilteredURL:   "/douyin/feed", // 设置你需要跳过认证的url,目前比较粗糙,是string而不是string数组,我们项目应该只需要一条URL吧
	})
	// 文件系统静态资源获取
	r.StaticFS("/cover", http.Dir("./cmd/api/static/images"))
	r.StaticFS("/videos", http.Dir("./cmd/api/static/videos"))
	v1 := r.Group("/douyin")
	v1.POST("/user/login/", authMiddleware.LoginHandler)
	v1.POST("/user/register/", handlers.Register, authMiddleware.LoginHandler) // 注册后自动登录
	// v1.GET("/feed", handlers.GetVideoFeed)
	//user1.Use(authMiddleware.MiddlewareFunc())
	v1.Use(authMiddleware.MiddlewareFunc())
	v1.GET("/feed", handlers.GetVideoFeed) // 有无登录正常写就行,未登陆的话claims为空,你查出来的userID是-1
	v1.GET("/user/", handlers.GetUserInfo)
	v1.GET("/publish/list", handlers.GetPublishVideoList)
	v1.POST("/publish/action/", handlers.PublishVideo)
	v1.POST("/favorite/action/", handlers.FavoriteByUser)
	v1.GET("/favorite/list/", handlers.GetFavoriteList)
	v1.POST("/comment/action/", handlers.PostComment)
	v1.GET("/comment/list/", handlers.QueryComments)
	// user2 := v1.Group("/relation")
	// v1.Use(authMiddleware.MiddlewareFunc())
	v1.GET("/relation/follow/list/", handlers.GetFollowList)
	v1.GET("/relation/follower/list/", handlers.GetFollowerList)
	v1.POST("/relation/action/", handlers.FollowAction) // 支持从postform里面获取token,不用加Bearer
	if err := http.ListenAndServe(":8080", r); err != nil {
		klog.Fatal(err)
	}
	r.Run()
}
