package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/yangjinguang/wechat-server/libs/xRes"
)

func HeaderParse() gin.HandlerFunc {
	return func(ct *gin.Context) {
		appId := ct.GetHeader("WeChat-AppId")
		if appId == "" {
			xRes.BadRequest(ct, "appId not found")
			ct.Abort()
		} else {
			ct.Set("appId", appId)
		}
		ct.Next()
	}
}
