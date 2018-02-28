package jwtToken

import (
	"github.com/dgrijalva/jwt-go"
	"time"
	"github.com/yangjinguang/wechat-server/libs/config"
)

type Token struct {
}

var jwtKey []byte

type UserClaims struct {
	Id       int64
	OpenId   string
	NickName string
	jwt.StandardClaims
}

func init() {
	if config.Conf.JwtKey != "" {
		jwtKey = []byte(config.Conf.JwtKey)
	} else {
		jwtKey = []byte("DVfK1X9z9qVXOg8IGZc2y3ThACrGQFJ1i1MsZ/xmwy9LSiHgmOQ12hpvVRX0ynePOzUvDhI")
	}
}

func (t *Token) New(userId int64, openId string, nickName string) (ss string, err error) {
	userClaims := UserClaims{
		userId,
		openId,
		nickName,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(7 * 24 * time.Hour).Unix(),
			Issuer:    "user",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims)
	ss, err = token.SignedString(jwtKey)
	return ss, err
}

func (t *Token) Parse(tokenString string) (userClaims UserClaims, err error) {
	token, err := jwt.ParseWithClaims(tokenString, &userClaims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		return *claims, nil
	} else {
		return userClaims, err
	}
}
