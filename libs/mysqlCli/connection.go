package mysqlCli

import (
	"fmt"
	"github.com/yangjinguang/wechat-server/libs/config"
	"github.com/yangjinguang/gomysql"
)

var Client *gomysql.DB

func init() {
	db, err := gomysql.Conn(fmt.Sprintf(
		"%s:%s@(%s:%d)/%s?parseTime=true&charset=utf8mb4,utf8",
		config.Conf.Mysql.User,
		config.Conf.Mysql.Password,
		config.Conf.Mysql.Host,
		config.Conf.Mysql.Port,
		config.Conf.Mysql.Database,
	))
	if err != nil {
		panic(err.Error())
	}
	Client = db
}

func Conn() *gomysql.DB {
	return Client.New()
}
