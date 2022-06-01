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
	//r := gin.New()
	//v1 := r.Group("/douyin")
	//video := v1.Group("/publish")
	//video.GET("/list", handlers.GetMyPublishVideoList)
	//if err := http.ListenAndServe(":8080", r); err != nil {
	//	klog.Fatal(err)
	//}
	r := gin.New()
	v1 := r.Group("/favorite")
	v1.GET("/action", handlers.FavoriteByUser)
	if err := http.ListenAndServe(":8080", r); err != nil {
		klog.Fatal(err)
	}

}
