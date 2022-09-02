package handle

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginHandle(c *gin.Context) {
	file := c.Param("file")
	c.File("./static/login/" + file)
}

func LoginPostHandle(c *gin.Context) {
	types := c.DefaultPostForm("type", "post")
	username := c.PostForm("username")
	password := c.PostForm("userpassword")
	c.String(http.StatusOK, fmt.Sprintf("username:%s,password:%s,type:%s", username, password, types))
}
