package dao

import (
	"WebServer/internal/dal/entity"
	"context"
	"fmt"

	"gorm.io/gorm"
)

type UserInfoDao struct {
	db *gorm.DB
}

func NewUserInfoDao(db *gorm.DB) *UserInfoDao {
	return &UserInfoDao{
		db: db,
	}
}

func (d *UserInfoDao) getDB(ctx context.Context) *gorm.DB {
	return d.db.Session(&gorm.Session{Context: ctx})
}

func (d *UserInfoDao) CreateUserInfo(ctx context.Context, user_info *entity.UserInfo) (err error) {
	db := d.getDB(ctx)
	if db == nil {
		return fmt.Errorf("[CreateUserInfo] empty db")
	}
	if err := db.Create(&user_info).Error; err != nil {
		return fmt.Errorf("[CreateUserInfo] db err: %w", err)
	}
	return nil
}
