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
	appSecret := ct.MustGet("appSecret").(string)
	type loginData struct {
		Code     string            `json:"code" binding:"required"`
		UserInfo models.WxUserInfo `json:"user_info" binding:"required"`
	}
	var ld loginData
	if err := ct.BindJSON(&ld); err == nil {
		ser := moduleUser.Service{}
		session, err := ser.GetSession(ld.Code, appId, appSecret)
		if err != nil {
			xRes.BadRequest(ct, err.Error())
			return
		}
		user, err := ser.Sync(appId, session.Openid, session.Unionid, ld.UserInfo)
		if err != nil {
			xRes.BadRequest(ct, err.Error())
			return
		}
		tk := jwtToken.Token{}
		token, err := tk.New(user.Id, user.OpenId, user.NickName)
		if err != nil {
			xRes.BadRequest(ct, err.Error())
			return
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
