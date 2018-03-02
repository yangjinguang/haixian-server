package moduleUser

import (
	"testing"
	"github.com/yangjinguang/wechat-server/libs/models"
)

var testId int64

func TestUser_Create(t *testing.T) {
	user := User{}
	user.NickName = "test-user"
	user.Gender = "1"
	id, err := user.Create()
	testId = id
	if err != nil {
		t.Error(err)
		t.Fail()
	} else {
		t.Log(id)
	}
}

func TestUser_GetAll(t *testing.T) {
	m := User{}
	users, err := m.GetAll()
	if err != nil {
		t.Error(err)
		t.Fail()
	} else {
		t.Log(users)
	}
}
func TestUser_GetById(t *testing.T) {
	m := &User{}
	user, _, err := m.GetById(testId)
	if err != nil {
		t.Error(err)
		t.Fail()
	} else {
		t.Log(user)
	}
}
func TestUser_Update(t *testing.T) {
	m := User{}
	user, _, err := m.GetById(testId)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	user.NickName = user.NickName + "-test-update"
	err = user.Update()
	if err != nil {
		t.Error(err)
		t.Fail()
	} else {
		t.Log("update success")
	}
}

func TestUser_DeleteById(t *testing.T) {
	m := User{}
	user, _, err := m.GetById(testId)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	err = user.DeleteById(18)
	if err != nil {
		t.Error(err)
		t.Fail()
	} else {
		t.Log("delete success")
	}
}

func TestService_Update(t *testing.T) {
	ser:=Service{}
	userInfo:=models.WxUserInfo{}
	userInfo.NickName = "test-user-2"
	user,err:=ser.Sync("test-app-2","test-open-id-2","",userInfo)
	if err != nil {
		t.Error(err)
		t.Fail()
	} else {
		t.Log(user)
		t.Log("update success")
	}
}