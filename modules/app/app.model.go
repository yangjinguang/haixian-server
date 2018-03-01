package moduleApp

import (
	"time"
	"fmt"
	"github.com/yangjinguang/wechat-server/libs/mysqlCli"
	"github.com/yangjinguang/wechat-server/libs/logger"
	"errors"
)

type App struct {
	Id        int64     `json:"id"`
	AppId     string    `json:"app_id"`
	AppSecret string    `json:"app_secret"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

const TABLE_NAME = "apps"

func (m *App) Insert() (id int64, err error) {
	db, err := mysqlCli.Conn()
	if err != nil {
		return id, err
	}
	sql := fmt.Sprintf(
		"insert into `%s` "+
			"(`app_id`,`app_secret`,`created_at`,`updated_at`) "+
			"values ('%s','%s','%s','%s')",
		TABLE_NAME,
		m.AppId,
		m.AppSecret,
		time.Now().Format("2006-01-02 15:04:05"),
		time.Now().Format("2006-01-02 15:04:05"),
	)
	logger.Debug(sql)
	result, err := db.Exec(sql)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (m *App) Replace() (err error) {
	db, err := mysqlCli.Conn()
	if err != nil {
		return err
	}
	if m.Id == 0 {
		return errors.New("invalid id")
	}
	sql := fmt.Sprintf(
		"replace into `%s` "+
			"(`id`,`app_id`,`app_secret`,`created_at`,`updated_at`) "+
			"values ('%d','%s','%s','%s','%s')",
		TABLE_NAME,
		m.Id,
		m.AppId,
		m.AppSecret,
		m.CreatedAt.Format("2006-01-02 15:04:05"),
		time.Now().Format("2006-01-02 15:04:05"),
	)
	_, err = db.Exec(sql)
	return err
}

func (m *App) Select(querySql string) (apps []*App, err error) {
	db, err := mysqlCli.Conn()
	if err != nil {
		return apps, err
	}
	sql := fmt.Sprintf("select `id`,`app_id`,`app_secret`,`created_at`,`updated_at` from `%s`", TABLE_NAME)
	if querySql != "" {
		sql += " " + querySql
	}
	logger.Debug(sql)
	rows, err := db.Query(sql)
	if err != nil {
		return apps, err
	}
	defer rows.Close()
	for rows.Next() {
		app := App{}
		err = rows.Scan(
			&app.Id,
			&app.AppId,
			&app.AppSecret,
			&app.CreatedAt,
			&app.UpdatedAt,
		)
		if err != nil {
			logger.Error(err)
			continue
		}
		apps = append(apps, &app)
	}
	return apps, nil
}

func (m *App) SelectOne(querySql string) (app *App, notFound bool, err error) {
	apps, err := m.Select(querySql)
	if err != nil {
		return app, false, err
	}
	if len(apps) <= 0 {
		return app, true, nil
	} else {
		return apps[0], false, nil
	}
}

func (m *App) Count(querySql string) (count int, err error) {
	db, err := mysqlCli.Conn()
	if err != nil {
		return 0, err
	}
	sql := fmt.Sprintf("select count(*) from `%s`", TABLE_NAME)
	if querySql != "" {
		sql += " " + querySql
	}
	rows, err := db.Query(sql)
	if err != nil {
		return 0, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			continue
		}
	}
	return count, nil
}

func (m *App) Delete(querySql string) (err error) {
	if querySql == "" {
		return errors.New("query sql is empty")
	}
	db, err := mysqlCli.Conn()
	if err != nil {
		return err
	}
	sql := fmt.Sprintf("delete from `%s` %s", TABLE_NAME, querySql)
	_, err = db.Exec(sql)
	return err
}
