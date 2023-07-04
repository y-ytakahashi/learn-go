package datastore

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMySqlDB() (conn *gorm.DB) {
	dsn := "docker:docker@tcp(127.0.0.1:4306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db, err := conn.DB()
	if err == nil {
		err = db.Ping()
		if err != nil {
			panic(err)
		}
	} else {
		panic(err)
	}
	return
}
