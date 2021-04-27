package middleware

import (
	"com.pippishen/trans-proxy/utils"
	"encoding/json"
	"fmt"
	"strconv"
	"testing"
)

func TestGenToken(t *testing.T)  {
	data := map[string]interface{}{
		"name":"david",
		"age": 23,
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
