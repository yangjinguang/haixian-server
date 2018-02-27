package mongoCli

import (
	"gopkg.in/mgo.v2"
	"log"
	"github.com/yangjinguang/wechat-server/libs/config"
)

var session *mgo.Session

func Conn() *mgo.Session {
	return session.Copy()
}

func init() {
	url := config.Conf.Mongo.Url
	s, err := mgo.Dial(url)
	if err != nil {
		log.Panic(err)
	}
	session = s
	session.SetMode(mgo.Monotonic, true)
}
