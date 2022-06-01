package main

import (
	"douyin/v1/cmd/api/handlers"
	"douyin/v1/cmd/api/rpc"
	"net/http"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gin-gonic/gin"
)

func init() {
	rpc.InitRpc()
}
func main() {
	r := gin.New()
	v1 := r.Group("/douyin")
	video := v1.Group("/publish")
	video.GET("/list", handlers.GetMyPublishVideoList)

	video_comments := v1.Group("/video_comments")
	video_comments.GET("/list", handlers.GetVideoCommentsList)
	video_comments.POST("", handlers.CommentAction)
	
	if err := http.ListenAndServe(":8080", r); err != nil {
		klog.Fatal(err)
	}
}
