package main

import (
	"TransProxy/enum"
	"TransProxy/manager"
	"TransProxy/manager/cache"
	"TransProxy/manager/db"
	"TransProxy/manager/mq"
	"TransProxy/model/request"
	"TransProxy/service"
	transPlatform "TransProxy/service/trans-platform"
	"TransProxy/service/translator"
	"TransProxy/utils"
	rand2 "crypto/rand"
	"fmt"
	"log"
	"math/big"
	"time"
)

func main() {
	// init manager
	service.InitManager()

	//release db
	if manager.TP_DB != nil {
		//main函数结束之前关闭资源
		defer db.Close()

		//初始化表和数据
		db.InitDB()
	}

	//release cache
	if manager.TP_CACHE_REDIS != nil {
		defer cache.Close()
	}

	//release mq
	if manager.TP_MQ_RABBIT != nil {
		defer mq.Close()
	}

	genData()
}

func genData() {
	google := &translator.Google{PlatformHandler: &transPlatform.SmartProxy{}}
	city := []string{
		"成都",
		"杭州",
		"重庆",
		"武汉",
		"西安",
		"苏州",
		"天津",
		"南京",
		"东莞",
		"郑州",
		"沈阳",
		"宁波",
		"昆明",
		"合肥",
	}

	for i := 0; i< 10000; i++ {
		randIndex, _  := rand2.Int(rand2.Reader, big.NewInt(int64(len(city))))
		randCity := city[randIndex.Int64()]

		UUID := utils.GetMD5Hash(utils.GetRandomString(10))
		item := request.Item {
			UUID: UUID,
			To: "en,ja",
			Platform: enum.Platform_Google,
			Text: randCity,
			Source: enum.Source_VAFFLE,
		}
		err := google.InsertItem(item)
		if err != nil {
			log.Println("Insert err: ", err)
		}

		fmt.Println("Insert item: ", item)
		time.Sleep(time.Millisecond * 300)
	}
}