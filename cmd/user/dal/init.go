package dal

import "douyin/v1/cmd/user/dal/db"

// Init init dal
func Init() {
	db.Init() // mysql
}
