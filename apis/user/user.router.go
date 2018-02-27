package apiUser

import "github.com/gin-gonic/gin"

func RouterInit(r *gin.RouterGroup) {
	c := Controller{}
	r.POST("/login", c.Login)
}
