package cache

import (
	"com.pippishen/trans-proxy/manager"
	TPTesting "com.pippishen/trans-proxy/manager/testing"
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
)

func TestConnect(t *testing.T)  {
	manager.TP_REDIS = Redis()
	convey.Convey("test connect redis", t, func() {
		convey.So(manager.TP_REDIS.Ping().String(), convey.ShouldEqual, "ping: PONG")
	})
}

func TestMain(m *testing.M) {
	fmt.Println("TestMain start")

	//初始化测试环境
	TPTesting.NEW().InitConfig()
	os.Exit(m.Run())
}