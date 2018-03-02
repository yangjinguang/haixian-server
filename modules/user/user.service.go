package moduleUser

import (
	"github.com/yangjinguang/wechat-server/libs/models"
	"github.com/yangjinguang/wechat-server/libs/xHttp"
	"encoding/json"
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

func (s *Service) GetSession(code string, appId string, appSecret string) (session models.WxUserSession, err error) {
	var query []models.KeyValue
	query = append(query, models.KeyValue{Key: "appid", Value: appId})
	query = append(query, models.KeyValue{Key: "secret", Value: appSecret})
	query = append(query, models.KeyValue{Key: "js_code", Value: code})
	query = append(query, models.KeyValue{Key: "grant_type", Value: "authorization_code"})
	_, body, err := xHttp.Get("https://api.weixin.qq.com/sns/jscode2session", query, []models.KeyValue{})
	if err != nil {
		return session, err
	}
	err = json.Unmarshal(body, &session)
	if err != nil {
		return session, err
	}
	return session, nil
}
