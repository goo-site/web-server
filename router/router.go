package router

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()
	r.POST("/form", func(c *gin.Context) {
		types := c.DefaultPostForm("type", "post")
		username := c.PostForm("username")
		password := c.PostForm("userpassword")
		c.String(http.StatusOK, fmt.Sprintf("username:%s,password:%s,type:%s", username, password, types))
	})
	r.GET("/", func(c *gin.Context) {
		c.File("./static/index.html")
	})

	if err := r.Run(":8080"); err != nil {
		fmt.Printf("startup service failed, err:%v\n", err)
	}
}
