package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormopentracing "gorm.io/plugin/opentracing"
	vc "douyin/v1/cmd/video_comments/dal/db"
)

var DB *gorm.DB

// Init init DB
func Init() {
	var err error
	DB, err = gorm.Open(mysql.Open("root:Yang75769933@tcp(139.224.195.12:3305)/video_comments?charset=utf8&parseTime=True&loc=Local"),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	println("数据库连接成功!")
	if err != nil {
		panic(err)
	}

	if err = DB.Use(gormopentracing.New()); err != nil {
		panic(err)
	}

	m := DB.Migrator()
	if m.HasTable(&vc.Video_Comments{}) {
		return
	}
	if err = m.CreateTable(&vc.Video_Comments{}); err != nil {
		panic(err)
	}
}