package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/yangjinguang/wechat-server/modules/user"
	"github.com/yangjinguang/wechat-server/libs/xRes"
)

func Authorized() gin.HandlerFunc {
	return func(ct *gin.Context) {
		token := ct.GetHeader("Authorization")
		if token == "" {
			xRes.Unauthorized(ct, "token not found")
		}
		user := moduleUser.User{}
		ct.Set("curUser", &user)
		ct.Next()
	}
}
