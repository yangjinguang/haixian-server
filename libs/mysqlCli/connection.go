package mysqlCli

import (
	"fmt"
	"github.com/yangjinguang/wechat-server/libs/config"
	"github.com/yangjinguang/wechat-server/libs/gomysql"
)

func Conn() (*gomysql.DB, error) {
	return gomysql.New(fmt.Sprintf(
		"%s:%s@(%s:%d)/%s?parseTime=true",
		config.Conf.Mysql.User,
		config.Conf.Mysql.Password,
		config.Conf.Mysql.Host,
		config.Conf.Mysql.Port,
		config.Conf.Mysql.Database,
	))

}
