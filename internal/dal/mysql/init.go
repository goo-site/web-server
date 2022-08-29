package mysql

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMySQL() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:12345678@(127.0.0.1:3306)/local?charset=utf8&parseTime=True&loc=Local&timeout=10s"))
	if err != nil {
		fmt.Println("open database error:", err)
	}

	sqlDB, _ := db.DB()
	sqlDB.SetConnMaxLifetime(300 * time.Second)
	sqlDB.SetMaxIdleConns(100)
	sqlDB.SetMaxOpenConns(200)

	return db
}
