package cache

import (
	"com.pippishen/trans-proxy/manager"
	TPTesting "com.pippishen/trans-proxy/manager/testing"
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
	"time"
)

func TestConnect(t *testing.T) {
	manager.TP_CACHE_REDIS = Redis()
	convey.Convey("test connect redis", t, func() {
		convey.So(manager.TP_CACHE_REDIS.Ping().String(), convey.ShouldEqual, "ping: PONG")
	})
}

func TestGetSet(t *testing.T) {
	manager.TP_CACHE_REDIS = Redis()
	convey.Convey("get and set value", t, func() {
		manager.TP_CACHE_REDIS.Set("username", "david", 10*time.Second)
		time.Sleep(8 * time.Second)
		v := manager.TP_CACHE_REDIS.Get("username")
		r, _ := v.Result()
		convey.So(r, convey.ShouldEqual, "david")
	})
}

func TestMain(m *testing.M) {
	fmt.Println("TestMain start")

	//初始化测试环境
	TPTesting.NEW().InitConfig()
	os.Exit(m.Run())
}
