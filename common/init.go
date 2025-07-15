package common

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var dsn = "root:@tcp(localhost:3306)/admin?charset=utf8mb4&parseTime=True&loc=Local"
var DB *gorm.DB

func init() {
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}
}
