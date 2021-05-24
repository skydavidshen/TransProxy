package config

import (
	"TransProxy/manager"
	"TransProxy/manager/log"
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
)

//application root dir
const rootDir = "/Users/davidshen/Documents/webroot/TransProxy"

func TestConfLoad(t *testing.T) {
	fmt.Println("TestConfLoad")

	//配置文件处理服务:支持热修改
	manager.TP_CONFIG = Viper()
	convey.Convey("get config for system.env_mode", t, func() {
		convey.So(manager.TP_SERVER_CONFIG.System.Env, convey.ShouldEqual, "dev")
	})
}

func TestConf(t *testing.T) {
	fmt.Println("env", manager.TP_SERVER_CONFIG.System.Oss)
}

func TestMain(m *testing.M) {
	fmt.Println("TestMain start")

	//初始化测试环境
	manager.TP_ROOT_DIR = rootDir
	manager.TP_CONFIG = Viper()
	//日志服务
	manager.TP_LOG = log.Zap()

	os.Exit(m.Run())
}