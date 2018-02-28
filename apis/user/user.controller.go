package apiUser

import (
	"github.com/gin-gonic/gin"
	"github.com/yangjinguang/wechat-server/libs/xRes"
	"github.com/yangjinguang/wechat-server/libs/models"
	"github.com/yangjinguang/wechat-server/modules/user"
	"github.com/yangjinguang/wechat-server/libs/jwtToken"
	"github.com/yangjinguang/wechat-server/libs/logger"
)

type Controller struct {
}

func (c *Controller) Login(ct *gin.Context) {
	appId := ct.MustGet("appId").(string)
	type loginData struct {
		OpenId   string            `json:"open_id" binding:"required"`
		UnionId  string            `json:"union_id"`
		UserInfo models.WxUserInfo `json:"user_info" binding:"required"`
	}
	var ld loginData
	if err := ct.BindJSON(&ld); err == nil {
		ser := moduleUser.Service{}
		user, err := ser.Update(appId, ld.OpenId, ld.UnionId, ld.UserInfo)
		if err != nil {
			xRes.BadRequest(ct, err.Error())
		}
		tk := jwtToken.Token{}
		token, err := tk.New(user.Id, user.OpenId, user.NickName)
		if err != nil {
			xRes.BadRequest(ct, err.Error())
		}
		xRes.OK(ct, gin.H{
			"user":  user,
			"token": token,
		})
	} else {
		xRes.MethodNotAllowed(ct, err.Error())
	}
}

func (c *Controller) Profile(ct *gin.Context) {
	curUser := ct.MustGet("curUser").(jwtToken.UserClaims)
	logger.Debug(curUser.NickName)
	xRes.OK(ct, gin.H{
		"user": curUser,
	})
}
