package apiExample

import (
	"github.com/gin-gonic/gin"
	"github.com/yangjinguang/wechat-server/libs/xUtils"
	"github.com/yangjinguang/wechat-server/libs/xRes"
)

type Controller struct {
}

func (c *Controller) Ping(ct *gin.Context) {
	user, err := xUtils.GetCurrentUser(ct)
	if err != nil {
		xRes.BadRequest(ct, err.Error())
	} else {
		xRes.OK(ct, user)
	}
}

func (c *Controller) Error(ct *gin.Context) {
	xRes.MethodNotAllowed(ct, nil)
}
