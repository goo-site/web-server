package mysql

import (
	"fmt"
	"time"
	"web-server/internal/dal/mysql/dao"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	UserInfoDao *dao.UserInfoDao
)

func openDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:12345678@(127.0.0.1:3306)/web_server_db?charset=utf8&parseTime=True&loc=Local&timeout=10s"))
	if err != nil {
		fmt.Println("open database error:", err)
	}

	sqlDB, _ := db.DB()
	sqlDB.SetConnMaxLifetime(300 * time.Second)
	sqlDB.SetMaxIdleConns(100)
	sqlDB.SetMaxOpenConns(200)

	return db
}

func InitMySQL() {
	db := openDB()

	UserInfoDao = dao.NewUserInfoDao(db)
}
