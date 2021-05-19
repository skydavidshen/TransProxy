package middleware

import (
	"TransProxy/utils"
	"encoding/json"
	"fmt"
	"strconv"
	"testing"
)

func TestGenToken(t *testing.T)  {
	data := map[string]interface{}{
		"uuid": "99a977b749fda07975953f52cff7e093",
		"platform": "google",
		"to": "en",
		"text": "你好",
		"source": "vaffle",
	}
	privateKey := "0YawE8IfRwHBzGQzo0cQD87B"
	timestamp := 1599653541

	dataJson, _ := json.Marshal(data)
	dataJsonStr := string(dataJson)
	// token算法: 对称hash加密
	// token = md5(md5(dataJsonStr) + privateKey + timestamp)
	preStr := fmt.Sprintf("%s%s%s", utils.GetMD5Hash(dataJsonStr), privateKey,
		strconv.Itoa(timestamp))
	genToken := utils.GetMD5Hash(preStr)
	utils.PrintObj(genToken)
}
