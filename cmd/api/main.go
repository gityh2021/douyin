package main

import (
	"context"
	"douyin/v1/cmd/api/rpc"
	"fmt"
)

func main()  {
	// 测试代码
	rpc.InitRpc()
	rpc.GetPublishVideoList(context.Background(), 1)
	fmt.Println(rpc.GetPublishVideoList(context.Background(), 1))
}