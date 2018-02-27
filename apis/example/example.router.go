package apiExample

import (
	"github.com/gin-gonic/gin"
)

func RouterInit(r *gin.RouterGroup) {
	c := Controller{}
	r.GET("/ping", c.Ping)
	r.GET("/error", c.Error)
}
