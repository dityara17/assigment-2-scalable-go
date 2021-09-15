package config

import (
	"assigment-2-scalable-go/model"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)


const DSN = "root@tcp(127.0.0.1:3306)/assignment_2_hackvtiv8?charset=utf8mb4&parseTime=True&loc=Local"

func DbInit() *gorm.DB {

	db, err := gorm.Open(mysql.Open(DSN), &gorm.Config{})

	if err != nil {
		panic("Can't connect to db")
	}
	err = db.AutoMigrate(model.Order{})
	if err != nil {
		panic(err.Error())
	}
	err = db.AutoMigrate(model.Items{})
	if err != nil {
		panic(err.Error())
	}
	return db
}
