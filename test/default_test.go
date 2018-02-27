package test

import (
	"testing"
	"os"
	"github.com/yangjinguang/wechat-server/libs/logger"
)
func init(){
	os.Setenv("LOG_LEVEL", "4")
}

func TestDefault(t *testing.T) {
	os.Setenv("LOG_SUFFIX", "Wechat-Server")
	logger.Debug(111)
}
