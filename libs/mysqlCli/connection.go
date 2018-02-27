package mysqlCli

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
	"github.com/yangjinguang/wechat-server/libs/config"
)

func Conn() (*sql.DB, error) {
	return sql.Open("mysql", fmt.Sprintf(
		"%s:%s@(%s:%d)/%s?parseTime=true",
		config.Conf.Mysql.User,
		config.Conf.Mysql.Password,
		config.Conf.Mysql.Host,
		config.Conf.Mysql.Port,
		config.Conf.Mysql.Database,
	))
}
