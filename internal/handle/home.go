package handle

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/goo-site/log"
	"net/http"
	"web-server/internal/dal/mysql"
)

func HomeHandle(c *gin.Context) {
	ctx := context.Background()
	userId := c.Request.Header.Get("user_id")
	userInfo, err := mysql.UserInfoDao.GetUserInfoByUserId(ctx, userId)
	if err != nil {
		log.Error("[HomeHandle] err:%v", err)
		c.JSON(http.StatusOK, gin.H{})
	}

	c.JSON(http.StatusOK, gin.H{
		"NickName": userInfo.Nickname,
	})
}
