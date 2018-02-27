package xUtils

import (
	"github.com/gin-gonic/gin"
	"errors"
	"github.com/yangjinguang/wechat-server/modules/user"
)

func GetCurrentUser(ct *gin.Context) (curUser *moduleUser.User, err error) {
	user, ok := ct.Get("curUser")
	if ok {
		curUser = user.(*moduleUser.User)
		return curUser, nil
	}
	return curUser, errors.New("current user not found")
}
