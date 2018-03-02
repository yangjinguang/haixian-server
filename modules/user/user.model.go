package moduleUser

import (
	"time"
	"github.com/yangjinguang/wechat-server/libs/mysqlCli"
	"fmt"
)

const TABLE_NAME = "users"

type User struct {
	Id        int64     `json:"id" mysql:"id"`
	UnionId   string    `json:"union_id" mysql:"union_id"`
	AppId     string    `json:"app_id" mysql:"app_id"`
	OpenId    string    `json:"open_id" mysql:"open_id"`
	NickName  string    `json:"nick_name" mysql:"nick_name"`
	Gender    int       `json:"gender" mysql:"gender"`
	Language  string    `json:"language" mysql:"language"`
	City      string    `json:"city" mysql:"city"`
	Province  string    `json:"province" mysql:"province"`
	Country   string    `json:"country" mysql:"country"`
	AvatarUrl string    `json:"avatar_url" mysql:"avatar_url"`
	CreatedAt time.Time `json:"created_at" mysql:"created_at"`
	UpdatedAt time.Time `json:"updated_at" mysql:"updated_at"`
}

func (m *User) Create() (int64, error) {
	db := mysqlCli.Conn()
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return db.T(TABLE_NAME).Insert(m)
}

func (m *User) Update() error {
	db := mysqlCli.Conn()
	m.UpdatedAt = time.Now()
	return db.T(TABLE_NAME).Replace(m)
}

func (m *User) GetAll() (users []*User, err error) {
	db := mysqlCli.Conn()
	err = db.T(TABLE_NAME).Select().All(&users)
	return users, err
}

func (m *User) GetById(id int64) (user *User, notFound bool, err error) {
	db := mysqlCli.Conn()
	user = &User{}
	notFound, err = db.T(TABLE_NAME).SelectById(id).One(user)
	return user, notFound, err
}

func (m *User) GetByAppIdAndOpenId(appId string, openId string) (user *User, notFound bool, err error) {
	db := mysqlCli.Conn()
	user = &User{}
	notFound, err = db.T(TABLE_NAME).Select().
		Where(fmt.Sprintf("`app_id` = '%s' and `open_id` = '%s'", appId, openId)).
		One(user)
	return user, notFound, err
}

func (m *User) DeleteById(id int64) error {
	db := mysqlCli.Conn()
	return db.T(TABLE_NAME).DeleteById(id)
}
