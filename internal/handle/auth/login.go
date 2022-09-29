package login

import (
	"fmt"
	"net/http"
	"web-server/internal/common/enum"

	"github.com/gin-gonic/gin"
	"web-server/internal/utils"
)

func LoginHandle(c *gin.Context) {
	types := c.DefaultPostForm("type", "post")
	username := c.PostForm("username")
	password := c.PostForm("userpassword")
	c.String(http.StatusOK, fmt.Sprintf("username:%s,password:%s,type:%s", username, password, types))
}

func RegisterHandle(c *gin.Context) {

}

func RetrieveHandle(c *gin.Context) {

}

func LoginAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		userId := utils.StringToInt64(c.Request.Header.Get("user_id"))

		switch CheckToken(token, userId) {
		case enum.TokenStatus_Invalid:

		}

	}
}
