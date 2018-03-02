package moduleApp

import (
	"time"
	"fmt"
	"github.com/yangjinguang/wechat-server/libs/mysqlCli"
)

type App struct {
	Id        int64     `json:"id" mysql:"id"`
	AppId     string    `json:"app_id" mysql:"app_id"`
	AppSecret string    `json:"app_secret" mysql:"app_secret"`
	CreatedAt time.Time `json:"created_at" mysql:"created_at"`
	UpdatedAt time.Time `json:"updated_at" mysql:"updated_at"`
}

const TABLE_NAME = "apps"

func (m *App) Create() (int64, error) {
	db := mysqlCli.Conn()
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return db.T(TABLE_NAME).Insert(m)
}

func (m *App) Update() error {
	db := mysqlCli.Conn()
	m.UpdatedAt = time.Now()
	return db.T(TABLE_NAME).Replace(m)
}

func (m *App) GetAll() (apps []*App, err error) {
	db := mysqlCli.Conn()
	err = db.T(TABLE_NAME).Select().All(&apps)
	return apps, err
}

func (m *App) GetById(id int64) (app *App, notFound bool, err error) {
	db := mysqlCli.Conn()
	app = &App{}
	notFound, err = db.T(TABLE_NAME).SelectById(id).One(app)
	return app, notFound, err
}

func (m *App) GetByAppId(appId string) (app *App, notFound bool, err error) {
	db := mysqlCli.Conn()
	app = &App{}
	notFound, err = db.T(TABLE_NAME).Select().
		Where(fmt.Sprintf("`app_id` = '%s'", appId)).
		One(app)
	return app, notFound, err
}

func (m *App) DeleteById(id int64) error {
	db := mysqlCli.Conn()
	return db.T(TABLE_NAME).DeleteById(id)
}
