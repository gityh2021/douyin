package main

import (
	"context"
	"net/http"
	"time"

	"douyin/v1/cmd/api/handlers"
	"douyin/v1/cmd/api/rpc"
	"douyin/v1/kitex_gen/user"
	"douyin/v1/pkg/constants"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gin-gonic/gin"
)

func Init() {
	rpc.InitRPC()
}

func main() {
	Init()
	r := gin.Default()
	authMiddleware, _ := jwt.New(&jwt.GinJWTMiddleware{
		Key:        []byte(constants.SecretKey),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			// fmt.Printf("data: %v\n", data)
			if v, ok := data.(int64); ok {
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

			return rpc.CheckUser(context.Background(), &user.CheckUserRequest{Username: loginVar1.UserName, Password: loginVar1.PassWord})
		},
		LoginResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			var loginVar2 handlers.UserParam
			loginVar2.UserName = c.Query("username")
			loginVar2.PassWord = c.Query("password")
			user_id, res := rpc.QueryUser(context.Background(), &user.CheckUserRequest{Username: loginVar2.UserName, Password: loginVar2.PassWord})
			handlers.SendLoginResponse(c, res, user_id, loginVar2.UserName, token, expire)
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})
	// 文件系统静态资源获取
	r.StaticFS("/cover", http.Dir("./cmd/api/static/images"))
	r.StaticFS("/videos", http.Dir("./cmd/api/static/videos"))
	v1 := r.Group("/douyin")
	// user1 := v1.Group("/user")
	v1.POST("/user/login/", authMiddleware.LoginHandler)
	v1.POST("/user/register/", handlers.Register, authMiddleware.LoginHandler) // 注册后自动登录
	v1.GET("/feed", handlers.GetVideoFeed)
	v1.POST("/publish/action/", func(c *gin.Context) {
		token := c.PostForm("token")
		c.Header("Authorization", token)
		c.Request.Header.Add("Authorization", token)
		//这里注意，看你是要加到c.Header还是c.Request.Header里，注释掉不要的一个即可
		//fmt.Println("c.GetHeader 的结果是 " + c.GetHeader("Authorization"))
		//fmt.Println("c.Request.Header.Get的结果是" + c.Request.Header.Get("Authorization"))
		newUrl := "/douyin/publish/action2/" //重定向的url
		c.Request.URL.Path = newUrl
		r.HandleContext(c)
	})
	//user1.Use(authMiddleware.MiddlewareFunc())
	v1.Use(authMiddleware.MiddlewareFunc())
	v1.GET("/user/", handlers.GetUserInfo)
	v1.GET("/publish/list", handlers.GetMyPublishVideoList)
	v1.POST("/publish/action2/", handlers.PublishVideo)
	v1.POST("/favorite/action/", handlers.FavoriteByUser)
	v1.GET("/favorite/list/", handlers.GetFavoriteLIst)
	v1.POST("/comment/action/", handlers.PostComment)
	v1.GET("/comment/list/", handlers.QueryComments)
	// user2 := v1.Group("/relation")
	// v1.Use(authMiddleware.MiddlewareFunc())
	v1.GET("/relation/follow/list/", handlers.GetFollowList)
	v1.GET("/relation/follower/list/", handlers.GetFollowerList)
	v1.POST("/relation/action/", handlers.FollowAction)
	if err := http.ListenAndServe(":8080", r); err != nil {
		klog.Fatal(err)
	}
	r.Run()
}
