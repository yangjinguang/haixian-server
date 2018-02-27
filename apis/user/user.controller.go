package apiUser

import (
	"github.com/gin-gonic/gin"
	"github.com/yangjinguang/wechat-server/libs/xRes"
	"github.com/yangjinguang/wechat-server/libs/logger"
)

type Controller struct {
}

func (c *Controller) Login(ct *gin.Context) {
	type loginData struct {
		UserName string `form:"username" json:"username" binding:"required"`
		Password string `form:"password" json:"password" binding:"required"`
	}
	var ld loginData
	if err := ct.BindJSON(&ld); err == nil {
		logger.Error(ld.UserName)
		xRes.OK(ct, gin.H{
			"login_data": ld,
		})
	} else {
		xRes.MethodNotAllowed(ct, err.Error())
	}
}
