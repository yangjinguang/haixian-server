package config

import (
	"log"
	"os/user"
	"path"
	"os"
	"encoding/json"
)

type Config struct {
	Port    string `json:"port"`
	HomeDir string `json:"home_dir"`
	TmpDir  string `json:"tmp_dir"`
	Mongo struct {
		Url string `json:"url"`
	} `json:"mongo"`
	Logger struct {
		Level  int    `json:"level"`
		File   string `json:"file"`
		Suffix string `json:"suffix"`
	} `json:"logger"`
	Mysql struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		Database string `json:"database"`
		User     string `json:"user"`
		Password string `json:"password"`
	} `json:"mysql"`
}

var Conf Config

func init() {
	usr, err := user.Current()
	if err != nil {
		log.Panic(err)
	}
	Conf.HomeDir = usr.HomeDir
	Conf.TmpDir = path.Join(usr.HomeDir, ".wechat-server")
	file, err := os.Open(Conf.TmpDir + "/config.json")
	defer file.Close()
	if err != nil {
		log.Panic(err)
	}
	jsonParser := json.NewDecoder(file)
	jsonParser.Decode(&Conf)
	if Conf.Port == "" {
		Conf.Port = "8080"
	}
}
