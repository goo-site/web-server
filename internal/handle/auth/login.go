package auth

import (
	"context"
	"fmt"
	"github.com/goo-site/log"
	"time"
	"web-server/internal/common/enum"
	"web-server/internal/dal/model"
	"web-server/internal/dal/mysql"
	"web-server/internal/handle/bo"

	"github.com/gin-gonic/gin"
)

func LoginHandle(ctx context.Context, req *bo.LoginRequest) error {
	userInfo, err := mysql.UserInfoDao.GetUserInfoByUserId(ctx, req.UserId)
	if err != nil {
		log.Error("[LoginHandle] GetUserInfo err:%v", err)
		return err
	}

	if userInfo.Password != req.PassWord {
		log.Error("[LoginHandle] Incorrect password")
		return fmt.Errorf("[LoginHandle] Incorrect password")
	}

	return nil
}

func RegisterHandle(ctx context.Context, req *bo.RegisterRequest) error {
	userInfo := &model.User{
		UserID:     req.UserId,
		Password:   req.PassWord,
		Email:      req.Email,
		Nickname:   req.UserId,
		Permission: 2,
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
	}

	err := mysql.UserInfoDao.CreateUserInfo(ctx, userInfo)
	if err != nil {
		log.Error("[RegisterHandle] CreateUserInfo err:%v", err)
		return err
	}
	return nil
}

func RetrieveHandle(c *gin.Context) {

}

func LoginAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		userId := c.Request.Header.Get("user_id")

		switch CheckToken(token, userId) {
		case enum.TokenStatus_Invalid, enum.TokenStatus_Expire:

		}
	}
}
