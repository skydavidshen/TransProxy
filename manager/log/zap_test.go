package log

import (
	"TransProxy/manager"
	TPTesting "TransProxy/manager/testing"
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"go.uber.org/zap"
	"os"
	"testing"
)

func TestLog(t *testing.T) {
	manager.TP_LOG = Zap()

	convey.Convey("fetch url", t, func() {
		url := "http://www.ddgooglesss.com"
		manager.TP_LOG.Error("Error fetch url..", zap.String("url", url))
	})
}

func TestMain(m *testing.M) {
	fmt.Println("TestMain start")

	//初始化测试环境
	TPTesting.NEW().InitConfig()
	os.Exit(m.Run())
}