package router

import (
	"fmt"
	"web-server/internal/handle"
	"web-server/internal/utils"

	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.New()
	r.Use(utils.Logger(), gin.Recovery())
	route(r)

	if err := r.Run(":8080"); err != nil {
		fmt.Printf("startup service failed, err:%v\n", err)
	}
}

func route(r *gin.Engine) {
	login := r.Group("/login")
	login.GET("/:file", handle.LoginHandle)
	login.POST("", handle.LoginPostHandle)
	r.GET("/", func(c *gin.Context) {
		c.Request.URL.Path = "/login/login.html"
		r.HandleContext(c)
	})
}
