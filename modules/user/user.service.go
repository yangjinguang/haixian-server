package moduleUser

import (
	"github.com/yangjinguang/wechat-server/libs/models"
)

type Service struct {
}

var m User

func (s *Service) Sync(appId string, openId string, unionId string, userInfo models.WxUserInfo) (newUser *User, err error) {
	fUser, notFound, err := m.GetByAppIdAndOpenId(appId, openId)
	if err != nil {
		return newUser, err
	}
	newUser = &User{}
	if notFound {
		newUser.AppId = appId
		newUser.OpenId = openId
		newUser.UnionId = unionId
		newUser.NickName = userInfo.NickName
		newUser.Gender = userInfo.Gender
		newUser.Language = userInfo.Language
		newUser.City = userInfo.City
		newUser.Province = userInfo.Province
		newUser.Country = userInfo.Country
		newUser.AvatarUrl = userInfo.AvatarUrl
		id, err := newUser.Create()
		if err != nil {
			return newUser, err
		}
		newUser.Id = id
	} else {
		newUser = fUser
		newUser.NickName = userInfo.NickName
		newUser.Gender = userInfo.Gender
		newUser.Language = userInfo.Language
		newUser.City = userInfo.City
		newUser.Province = userInfo.Province
		newUser.Country = userInfo.Country
		newUser.AvatarUrl = userInfo.AvatarUrl
		err := newUser.Update()
		if err != nil {
			return newUser, err
		}
	}
	return newUser, err
}
