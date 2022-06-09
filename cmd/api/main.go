package main

import (
	"context"
	"douyin/v1/cmd/api/oss"
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
	oss.Init()
}

func main() {
	Init()
	go oss.ConsumeImg()
	go oss.ConsumeVideo()
	r := gin.Default()

	//jwt的中间件配置
	authMiddleware, _ := myjwt.New(&myjwt.GinJWTMiddleware{
		Key:        []byte(constants.SecretKey), //服务端密钥
		Timeout:    time.Hour,                   //token有效期为1个小时
		MaxRefresh: time.Hour,                   //token更新时间为1个小时

		//调用顺序 Authenticator->PayloadFunc->LoginResponse
		//成功登陆后使用 生成令牌
		PayloadFunc: func(data interface{}) myjwt.MapClaims {
			if v, ok := data.(int64); ok {
				return myjwt.MapClaims{
					constants.IdentityKey: v,
				}
			}
			return myjwt.MapClaims{}
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"status_code": int32(code),
				"status_msg":  message,
				"user_id":     0,
				"token":       "",
			})
		},

		//Authenticator 登陆信息的校验 返回uid和error
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVar1 handlers.UserParam
			//这里不用Query而是用ShouldBind的话重复登陆会有问题
			loginVar1.UserName = c.Query("username")
			loginVar1.PassWord = c.Query("password")
			if len(loginVar1.UserName) == 0 || len(loginVar1.PassWord) == 0 {
				return "", myjwt.ErrMissingLoginValues
			}

			return rpc.CheckUser(context.Background(), &user.CheckUserRequest{Username: loginVar1.UserName, Password: loginVar1.PassWord})
		},

		//LoginResponse 将生成的令牌作为JSON数据返回用户
		LoginResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			var loginVar2 handlers.UserParam
			loginVar2.UserName = c.Query("username")
			loginVar2.PassWord = c.Query("password")
			user_id, res := rpc.QueryUser(context.Background(), &user.CheckUserRequest{Username: loginVar2.UserName, Password: loginVar2.PassWord})
			handlers.SendLoginResponse(c, res, user_id, loginVar2.UserName, token, expire)
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt, postform: token",                                                                                 //设置jwt的获取位置
		TokenHeadName: "Bearer",                                                                                                                                            //Header中的token头部字段 默认bearer即可
		TimeFunc:      time.Now,                                                                                                                                            //设置时间函数
		FilteredURL:   "/douyin/feed, /douyin/publish/list, /douyin/favorite/list/, /douyin/comment/list/, /douyin/relation/follow/list/, /douyin/relation/follower/list/", // 设置你需要跳过认证的url
	})
	v1 := r.Group("/douyin")
	v1.POST("/user/login/", authMiddleware.LoginHandler)                       //登陆操作
	v1.POST("/user/register/", handlers.Register, authMiddleware.LoginHandler) // 注册后自动登录
	v1.Use(authMiddleware.MiddlewareFunc())                                    //这里下面的开启jwt认证
	v1.GET("/feed", handlers.GetVideoFeed)                                     // 获取视频流有无登录正常写就行,未登陆的话claims为空,你查出来的userID是-1
	v1.GET("/user/", handlers.GetUserInfo)                                     //获取用户信息
	v1.GET("/publish/list", handlers.GetPublishVideoList)                      //获取发布视频的列表
	v1.POST("/publish/action/", handlers.PublishVideo)                         //发布视频
	v1.POST("/favorite/action/", handlers.FavoriteByUser)                      //点赞视频
	v1.GET("/favorite/list/", handlers.GetFavoriteList)                        //获取点赞列表
	v1.POST("/comment/action/", handlers.PostComment)                          //进行评论
	v1.GET("/comment/list/", handlers.QueryComments)                           //获取评论列表
	v1.GET("/relation/follow/list/", handlers.GetFollowList)                   //获取关注列表
	v1.GET("/relation/follower/list/", handlers.GetFollowerList)               //获取粉丝列表
	v1.POST("/relation/action/", handlers.FollowAction)                        // 进行关注或者取关，支持从postform里面获取token,不用加Bearer
	if err := http.ListenAndServe(":"+constants.API_PORT, r); err != nil {
		klog.Fatal(err)
	}
	r.Run()
}
