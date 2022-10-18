package router

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"web-server/internal/common/enum"
	"web-server/internal/handle"
	"web-server/internal/handle/auth"
	"web-server/internal/handle/bo"
	"web-server/internal/utils"
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
	// 权限路由组
	l := r.Group("/auth")
	l.POST("/login", func(c *gin.Context) {
		var req bo.LoginRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, bo.NewErrorResp("parse LoginRequest err"))
			return
		}

		ctx := context.Background()
		err := auth.LoginHandle(ctx, &req)
		if err != nil {
			c.JSON(http.StatusBadRequest, bo.NewErrorResp(err.Error()))
			return
		}

		token, err := auth.CreateToken(req.UserId)
		if err != nil {
			c.JSON(http.StatusBadRequest, bo.NewErrorResp(err.Error()))
			return
		}

		resp := &bo.Resp{
			Status: enum.RespStatus_OK,
		}

		c.Header("token", token)
		c.JSON(http.StatusOK, bo.NewSuccessResp(utils.MarshalString(resp)))
	})

	l.POST("/register", func(c *gin.Context) {
		var req *bo.RegisterRequest
		if err := c.ShouldBindJSON(req); err != nil {
			c.JSON(http.StatusBadRequest, bo.NewErrorResp("parse RegisterRequest err"))
			return
		}

		ctx := context.Background()
		err := auth.RegisterHandle(ctx, req)
		if err != nil {
			c.JSON(http.StatusBadRequest, bo.NewErrorResp(err.Error()))
			return
		}

		resp := &bo.Resp{
			Status: enum.RespStatus_OK,
		}

		c.JSON(http.StatusOK, bo.NewSuccessResp(utils.MarshalString(resp)))
	})

	l.POST("/retrieve", auth.RetrieveHandle)

	// 主页路由组
	home := r.Group("/home", auth.LoginAuth())
	home.GET("", handle.HomeHandle)
}
