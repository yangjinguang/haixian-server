package moduleUser

type AppUser struct {
	Id        int    `json:"id"`
	AppId     string `json:"app_id"`
	UserId    int    `json:"user_id"`
	OpenId    string `json:"open_id"`
	UnionId   string `json:"union_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (m *AppUser) GetAll() (appUsers []AppUser, err error) {
	return appUsers, nil
}

func (m *AppUser) GetById(id int) (appUser AppUser, err error) {
	return appUser, nil
}

func (m *AppUser) GetByUnionId(unionId string) (appUsers []AppUser, err error) {
	return appUsers, nil
}

func (m *AppUser) GetByUserId(UserId int) (appUsers []AppUser, err error) {
	return appUsers, nil
}

func (m *AppUser) GetByAppId(appId string) (appUsers []AppUser, err error) {
	return appUsers, nil
}

func (m *AppUser) GetByOpenId(openId string) (appUsers []AppUser, err error) {
	return appUsers, nil
}

func (m *AppUser) Save() (id int, err error) {
	return id, nil
}

func (m *AppUser) UpdateById(id int) (err error) {
	return nil
}

func (m *AppUser) DeleteById(id int) (err error) {
	return nil
}
