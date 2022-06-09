package db

import (
	"douyin/v1/pkg/constants"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"gorm.io/plugin/dbresolver"
)

var DB *gorm.DB

// Init 初始化DB
func Init() {
	var err error
	DB, err = gorm.Open(mysql.Open(constants.MySQLDefaultDSN),
		&gorm.Config{
			PrepareStmt: true,
		},
	)
	if err != nil {
		panic(err)
	}
	err = DB.Use(dbresolver.Register(dbresolver.Config{
		Sources:  []gorm.Dialector{mysql.Open(constants.MySQLDefaultDSN)},
		Replicas: []gorm.Dialector{mysql.Open(constants.MySQLReplicaDSN)},
		// sources/replicas 负载均衡策略
		Policy: dbresolver.RandomPolicy{},
	}))
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
