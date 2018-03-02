package models

type WxUserInfo struct {
	NickName  string `json:"nickName"`
	Gender    int    `json:"gender"`
	Language  string `json:"language"`
	City      string `json:"city"`
	Province  string `json:"province"`
	Country   string `json:"country"`
	AvatarUrl string `json:"avatarUrl"`
}

type KeyValue struct {
	Key   string `json:"key" bson:"key"`
	Value string `json:"value" bson:"value"`
}

type WxUserSession struct {
	Openid     string `json:"openid"`
	SessionKey string `json:"session_key"`
	Unionid    string `json:"unionid"`
}