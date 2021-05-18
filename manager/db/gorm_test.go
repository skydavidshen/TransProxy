package db

import (
	"TransProxy/manager"
	TPTesting "TransProxy/manager/testing"
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"gorm.io/gorm"
	"math/rand"
	"os"
	"testing"
)

type User struct {
	gorm.Model
	Name string `json:"name" gorm:"type:varchar(80);not null;unique;comment:用户登录名;default:'';"`
	Age  int    `json:"age" gorm:"type:int(4);not null;comment:用户年龄;default:1;"`
}

func TestConn(t *testing.T) {
	manager.TP_DB = Gorm()
	defer Close()

	convey.Convey("auto migrate", t, func() {
		err := manager.TP_DB.AutoMigrate(&User{})
		convey.So(err, convey.ShouldEqual, nil)
	})
}

func TestInsert(t *testing.T) {
	manager.TP_DB = Gorm()
	defer Close()

	convey.Convey("insert user data", t, func() {
		user := &User{
			Name: fmt.Sprintf("davids%d", rand.Int()),
			Age:  rand.Intn(100),
		}
		manager.TP_DB.Debug().Create(user)
		convey.So(user.ID, convey.ShouldBeGreaterThan, 1)
	})
}

func TestSelect(t *testing.T) {
	manager.TP_DB = Gorm()
	defer Close()

	convey.Convey("select item from user table", t, func() {
		var user User
		manager.TP_DB.Debug().First(&user)
		fmt.Printf("user: %v", user)
		convey.So(user.Name, convey.ShouldEqual, "david")
	})
}

func TestMain(m *testing.M) {
	fmt.Println("TestMain start")

	//初始化测试环境
	TPTesting.NEW().InitConfig()
	os.Exit(m.Run())
}
