package main

import (
	"douyin/v1/cmd/video/dal/db"
	"douyin/v1/kitex_gen/video/videoservice"
	"douyin/v1/pkg/constants"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/registry-etcd"
	"net"
)

func main() {
	r, err := etcd.NewEtcdRegistry([]string{"139.224.195.12:2379"}) // r should not be reused.
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", ":"+constants.VIDEO_PORT)
	if err != nil {
		panic(err)
	}
	db.Init()
	svr := videoservice.NewServer(new(VideoServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "video"}), // server name
		server.WithServiceAddr(addr),                                       // address
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}), // limit
		server.WithMuxTransport(),                                          // Multiplex
		//server.WithBoundHandler(bound.NewCpuLimitHandler()),                // BoundHandler
		server.WithRegistry(r), // registry
	)
	err = svr.Run()
	if err != nil {
		klog.Fatal(err)
	}
}
