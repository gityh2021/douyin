package main

import (
	"douyin/v1/cmd/api/handlers"
	"douyin/v1/cmd/api/rpc"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {
	rpc.InitRpc()
}
func main() {
	r := gin.Default()
	// 文件系统静态资源获取
	r.StaticFS("/cover", http.Dir("./cmd/api/static/images"))
	r.StaticFS("/videos", http.Dir("./cmd/api/static/videos"))
	//r.Static("/video", "./static/videos")
	v1 := r.Group("/douyin")
	v1.GET("/feed", handlers.GetVideoFeed)
	v1.GET("/publish/list", handlers.GetMyPublishVideoList)
	v1.POST("/publish/action/", handlers.PublishVideo)
	if err := http.ListenAndServe(":8080", r); err != nil {
		klog.Fatal(err)
	}
	r.Run()
}
