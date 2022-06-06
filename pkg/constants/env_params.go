package constants

import "os"

var ServerAddress = os.Getenv("SERVER_IP")
var NetworkAddress = os.Getenv("NETWORK_IP")
var USER_PORT = os.Getenv("USER_PORT")
var API_PORT = os.Getenv("API_PORT")
var VIDEO_PORT = os.Getenv("VIDEO_PORT")
var PlayURL = "http://" + ServerAddress + ":" + API_PORT + "/videos/"
var CoverURL = "http://" + ServerAddress + ":" + API_PORT + "/cover/"
