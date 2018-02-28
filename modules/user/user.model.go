package moduleUser

import (
	"github.com/yangjinguang/wechat-server/libs/mysqlCli"
	"fmt"
	"errors"
	"time"
	"github.com/yangjinguang/wechat-server/libs/logger"
)

const TABLE_NAME = "users"

type User struct {
	Id        int64     `json:"id"`
	UnionId   string    `json:"union_id"`
	AppId     string    `json:"app_id"`
	OpenId    string    `json:"open_id"`
	NickName  string    `json:"nick_name"`
	Gender    string    `json:"gender"`
	Language  string    `json:"language"`
	City      string    `json:"city"`
	Province  string    `json:"province"`
	Country   string    `json:"country"`
	AvatarUrl string    `json:"avatar_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (m *User) GetAll() (users []*User, err error) {
	return m.DbSelect("")
}

func (m *User) GetById(id int64) (user *User, notFound bool, err error) {
	users, err := m.DbSelect(fmt.Sprintf("where `id` = '%d'", id))
	if err != nil {
		return user, false, err
	}
	if len(users) <= 0 {
		return user, true, nil
	} else {
		return users[0], false, nil
	}
}

func (m *User) GetByUnionId(unionId string) (user *User, notFound bool, err error) {
	users, err := m.DbSelect(fmt.Sprintf("where `UnionId` = '%s'", unionId))
	if err != nil {
		return user, false, err
	}
	if len(users) <= 0 {
		return user, true,nil
	} else {
		return users[0], false, nil
	}
}

func (m *User) GetByAppIdAndOpenId(appId string, openId string) (user *User, notFound bool, err error) {
	users, err := m.DbSelect(fmt.Sprintf("where `app_id` = '%s' and `open_id` = '%s'", appId, openId))
	if err != nil {
		return user, false, err
	}
	if len(users) <= 0 {
		return user, true, nil
	} else {
		return users[0], false, nil
	}
}

func (m *User) Create() (id int64, err error) {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return m.DbInsert()
}

func (m *User) Update() (err error) {
	m.UpdatedAt = time.Now()
	return m.DbReplace()
}

func (m *User) DeleteById() (err error) {
	return m.DbDelete(fmt.Sprintf("where `id` = '%d'", m.Id))
}

func (m *User) DbInsert() (id int64, err error) {
	db, err := mysqlCli.Conn()
	if err != nil {
		return id, err
	}
	sql := fmt.Sprintf(
		"insert into `%s` "+
			"(`union_id`,`app_id`,`open_id`,`nick_name`,`gender`,`language`,`city`,`province`,`country`,`avatar_url`,`created_at`,`updated_at`) "+
			"values ('%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s')",
		TABLE_NAME,
		m.UnionId,
		m.AppId,
		m.OpenId,
		m.NickName,
		m.Gender,
		m.Language,
		m.City,
		m.Province,
		m.Country,
		m.AvatarUrl,
		m.CreatedAt.Format("2006-01-02 15:04:05"),
		m.UpdatedAt.Format("2006-01-02 15:04:05"),
	)
	logger.Debug(sql)
	result, err := db.Exec(sql)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (m *User) DbReplace() (err error) {
	db, err := mysqlCli.Conn()
	if err != nil {
		return err
	}
	if m.Id == 0 {
		return errors.New("not found")
	}
	sql := fmt.Sprintf(
		"replace into `%s` "+
			"(`id`,`union_id`,`app_id`,`open_id`,`nick_name`,`gender`,`language`,`city`,`province`,`country`,`avatar_url`,`created_at`,`updated_at`) "+
			"values ('%d','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s')",
		TABLE_NAME,
		m.Id,
		m.UnionId,
		m.AppId,
		m.OpenId,
		m.NickName,
		m.Gender,
		m.Language,
		m.City,
		m.Province,
		m.Country,
		m.AvatarUrl,
		m.CreatedAt.Format("2006-01-02 15:04:05"),
		m.UpdatedAt.Format("2006-01-02 15:04:05"),
	)
	_, err = db.Exec(sql)
	return err
}

func (m *User) DbSelect(querySql string) (users []*User, err error) {
	db, err := mysqlCli.Conn()
	if err != nil {
		return users, err
	}
	sql := fmt.Sprintf("select `id`,`union_id`,`app_id`,`open_id`,`nick_name`,`gender`,`language`,`city`,`province`,`country`,`avatar_url`,`created_at`,`updated_at` from `%s`", TABLE_NAME)
	if querySql != "" {
		sql += " " + querySql
	}
	logger.Debug(sql)
	rows, err := db.Query(sql)
	if err != nil {
		return users, err
	}
	defer rows.Close()
	for rows.Next() {
		user := User{}
		err = rows.Scan(
			&user.Id,
			&user.UnionId,
			&user.AppId,
			&user.OpenId,
			&user.NickName,
			&user.Gender,
			&user.Language,
			&user.City,
			&user.Province,
			&user.Country,
			&user.AvatarUrl,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			logger.Error(err)
			continue
		}
		users = append(users, &user)
	}
	return users, nil
}

func (m *User) DbCount(querySql string) (count int, err error) {
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

func (m *User) DbDelete(querySql string) (err error) {
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
