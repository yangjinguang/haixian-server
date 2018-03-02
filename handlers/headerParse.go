package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/yangjinguang/wechat-server/libs/xRes"
	"github.com/yangjinguang/wechat-server/modules/app"
)

func HeaderParse() gin.HandlerFunc {
	return func(ct *gin.Context) {
		appId := ct.GetHeader("WeChat-AppId")
		if appId == "" {
			xRes.BadRequest(ct, "appId not found")
			ct.Abort()
		} else {
			appM := moduleApp.App{}
			app, notFound, err := appM.GetByAppId(appId)
			if err != nil {
				xRes.BadRequest(ct, err.Error())
				ct.Abort()
			}
			if notFound {
				xRes.BadRequest(ct, "appId is invalid")
				ct.Abort()
			}
			ct.Set("appId", app.AppId)
			ct.Set("appSecret", app.AppSecret)
		}
		ct.Next()
	}
}
