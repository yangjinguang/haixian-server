package models

type WxUserInfo struct {
	NickName  string `json:"nickName"`
	Gender    string `json:"gender"`
	Language  string `json:"language"`
	City      string `json:"city"`
	Province  string `json:"province"`
	Country   string `json:"country"`
	AvatarUrl string `json:"avatarUrl"`
}
