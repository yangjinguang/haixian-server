package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/yangjinguang/wechat-server/modules/user"
)

func Authorized() gin.HandlerFunc {
	return func(ct *gin.Context) {
		user := moduleUser.User{}
		ct.Set("curUser", &user)
		ct.Next()
	}
}
