package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

func init() {
	var err error
	dns := "root:123456@tcp(127.0.0.1:3306)/document_server?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		//panic("failed to connect to database")
		log.Fatalf("Failed to connect to database, err: %v", err)
	}

}
