package daemon

import (
	"TransProxy/enum"
	"TransProxy/model/request"
	transPlatform "TransProxy/service/trans-platform"
	"TransProxy/service/translator"
	"TransProxy/utils"
	rand2 "crypto/rand"
	"fmt"
	"log"
	"math/big"
	"time"
)

type GenTestingData struct {}

func (g GenTestingData) DoTask() {
	fmt.Println("Do task: GenTestingData...")

	genData(10000)
}

func genData(count int) {
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

	for i := 0; i< count; i++ {
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