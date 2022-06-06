package db

import (
	"douyin/v1/pkg/constants"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Init init DB
func Init() {
	var err error
	DB, err = gorm.Open(mysql.Open(constants.MySQLDefaultDSN),
		&gorm.Config{
			PrepareStmt: true,
			//SkipDefaultTransaction: true,
			//这里要注释掉 不然不会开启事务
		},
	)
	if err != nil {
		panic(err)
	}

	// 设为手动迁移，通过Hash看一下是不是有重复的表
	// user表
	m := DB.Migrator()
	if m.HasTable(&User{}) {
		return
	}
	if err = m.CreateTable(&User{}); err != nil {
		panic(err)
	}
	// follower表
	if m.HasTable(&Follower{}) {
		return
	}
	if err = m.CreateTable(&Follower{}); err != nil {
		panic(err)
	}
}
