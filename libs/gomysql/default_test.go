package gomysql

import (
	"testing"
	"fmt"
	"github.com/yangjinguang/wechat-server/libs/config"
	"github.com/yangjinguang/wechat-server/libs/logger"
	"encoding/json"
	"time"
)

var db *DB

type TmpUser struct {
	Id        int64     `json:"id" mysql:"id"`
	NickName  string    `json:"nick_name" mysql:"nick_name"`
	CreatedAt time.Time `json:"created_at" mysql:"created_at"`
	UpdatedAt time.Time `json:"updated_at" mysql:"updated_at"`
}

func init() {
	d, err := New(fmt.Sprintf(
		"%s:%s@(%s:%d)/%s?parseTime=true",
		config.Conf.Mysql.User,
		config.Conf.Mysql.Password,
		config.Conf.Mysql.Host,
		config.Conf.Mysql.Port,
		config.Conf.Mysql.Database,
	))
	db = d
	if err != nil {
		panic(err.Error())
	}
}

func TestDB_All(t *testing.T) {
	var users []*TmpUser

	db.T("users").
		Select().
		Where("").
		Limit(0, 1).
		All(&users)
	s, _ := json.Marshal(users)
	logger.Info(string(s))
}

func TestDB_One(t *testing.T) {
	user := TmpUser{}
	_,err := db.T("users").
		SelectById(88).
		One(&user)
	logger.Debug(err)
	logger.Debug(user)
}

func TestDB_Insert(t *testing.T) {
	user := TmpUser{}
	user.NickName = "test-444"
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	id, _ := db.T("users").Insert(&user)
	logger.Debug(id)
}

func TestDB_Replace(t *testing.T) {
	user := TmpUser{}
	user.Id = 1
	user.NickName = "test-444"
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	db.T("users").Replace(&user)
}

func TestDB_Delete(t *testing.T) {
	err := db.T("users").Where("`nick_name` = 'test-444'").Delete()
	logger.Debug(err)
}
