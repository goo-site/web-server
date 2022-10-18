package dao

import (
	"context"
	"fmt"
	"github.com/goo-site/log"
	"web-server/internal/dal/model"

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

func (d *UserInfoDao) CreateUserInfo(ctx context.Context, userInfo *model.User) error {
	db := d.getDB(ctx)
	if db == nil {
		log.Error("[CreateUserInfo] empty db")
		return fmt.Errorf("[CreateUserInfo] empty db")
	}
	if err := db.Create(&userInfo).Error; err != nil {
		log.Error("[CreateUserInfo] db err: %w", err)
		return err
	}
	return nil
}

func (d *UserInfoDao) GetUserInfoByUserId(ctx context.Context, userId string) (*model.User, error) {
	db := d.getDB(ctx)
	if db == nil {
		log.Error("[GetUserInfo] empty db")
		return nil, fmt.Errorf("[GetUserInfo] empty db")
	}

	userInfo := &model.User{}
	if err := db.Where("user_id = ?", userId).First(&userInfo).Error; err != nil {
		return nil, fmt.Errorf("[GetUserInfo] db err: %w", err)
	}
	return userInfo, nil
}
