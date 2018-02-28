package apiUser

import (
	"github.com/gin-gonic/gin"
	"github.com/yangjinguang/wechat-server/handlers"
)

func RouterInit(r *gin.RouterGroup) {
	c := Controller{}
	r.POST("/login", c.Login)
	r.GET("/profile", handlers.Authorized(),c.Profile)
}
