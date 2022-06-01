package repository

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormopentracing "gorm.io/plugin/opentracing"
)

var DB *gorm.DB

// Init init DB
func Init() {
	var err error
	DB, err = gorm.Open(mysql.Open("root:Yang75769933@tcp(139.224.195.12:3305)/video?charset=utf8&parseTime=True&loc=Local"),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	println("数据库连接成功!")
	if err != nil {
		panic(err)
	}
	err = Register()
	if err != nil {
		return
	}
	err = CreateTable()
	if err != nil {
		return
	}

}
func CreateUser() error {
	m := DB.Migrator()
	if m.HasTable(&User{}) {
		return nil
	}
	if err := m.CreateTable(&User{}); err != nil {
		return err
	}
	return nil
}
func CreateTable() error {
	err := CreateUser()
	if err != nil {
		return err
	}
	return nil
}
func Register() error {
	if err := DB.Use(gormopentracing.New()); err != nil {
		return err
	}
	return nil
}
