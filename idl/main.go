package main

import (
	"log"

	userdemo "github.com/Baojiazhong/dousheng-ubuntu/kitex_gen/userdemo/userservice"
)

func main() {
	svr := userdemo.NewServer(new(UserServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
	// TODO: Your code here...
	// 中文注释
}
