package constants

import "os"

var ServerAddress = os.Getenv("SERVER_IP")
var NetworkAddress = os.Getenv("NETWORK_IP")
var PlayURL = "http://" + ServerAddress + ":8080/videos/"
var CoverURL = "http://" + ServerAddress + ":8080/cover/"
