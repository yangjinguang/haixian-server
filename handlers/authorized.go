package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/yangjinguang/wechat-server/libs/xRes"
	"github.com/yangjinguang/wechat-server/libs/jwtToken"
)

func Authorized() gin.HandlerFunc {
	return func(ct *gin.Context) {
		token := ct.GetHeader("Authorization")
		if token == "" {
			xRes.Unauthorized(ct, "token not found")
		}
		tk := jwtToken.Token{}
		userClaims, err := tk.Parse(token)
		if err != nil {
			xRes.Unauthorized(ct, err.Error())
		}
		ct.Set("curUser", userClaims)
		ct.Next()
	}
}
