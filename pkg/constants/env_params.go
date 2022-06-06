package constants

import "os"

// 部署环境
var NetworkAddress = os.Getenv("NETWORK_IP")
var USER_PORT = os.Getenv("USER_PORT")
var VIDEO_PORT = os.Getenv("VIDEO_PORT")
var API_PORT = os.Getenv("API_PORT")

// Debug环境
//var NetworkAddress = "127.0.0.1"
//var USER_PORT = "8087"
//var VIDEO_PORT = "8088"
//var API_PORT = "8082"
