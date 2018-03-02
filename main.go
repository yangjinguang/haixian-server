package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/yangjinguang/wechat-server/handlers"
	"github.com/yangjinguang/wechat-server/apis/example"
	"github.com/yangjinguang/wechat-server/apis/user"
	"github.com/yangjinguang/wechat-server/libs/config"
)

func main() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(handlers.HeaderParse())
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "it's worked !!",
		})
	})
	apiExample.RouterInit(r.Group("/api/v1", handlers.Authorized()))
	apiUser.RouterInit(r.Group("/api/user"))
	port := config.Conf.Port
	r.Run(":" + port) // listen and serve on 0.0.0.0:8080
}
