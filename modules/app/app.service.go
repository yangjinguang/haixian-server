package moduleApp

import (
	"fmt"
	"time"
)

type Service struct {
}

var m App

func (s *Service) GetAll() (apps []*App, err error) {
	return m.DbSelect("")
}

func (s *Service) GetById(id int64) (*App, bool, error) {
	return m.DbSelectOne(fmt.Sprintf("where `id` = '%d'", id))
}

func (s *Service) GetByAppId(appId string) (*App, bool, error) {
	return m.DbSelectOne(fmt.Sprintf("where `app_id` = '%s'", appId))
}

func (s *Service) Create() (id int64, err error) {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return m.DbInsert()
}

func (s *Service) Update() (err error) {
	m.UpdatedAt = time.Now()
	return m.DbReplace()
}

func (s *Service) DeleteById(id int64) (err error) {
	return m.DbDelete(fmt.Sprintf("where `id` = '%d'", m.Id))
}