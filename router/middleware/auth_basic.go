package middleware

import (
	"TransProxy/manager"
	"TransProxy/model/request"
	"TransProxy/model/response"
	"TransProxy/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io/ioutil"
	"strconv"
)

func AuthBasic() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 重新写回Request Body, 供controller使用
		bodyBytes, _ := ioutil.ReadAll(c.Request.Body)
		c.Request.Body.Close() //  must close
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

		var basic request.Basic
		_ = json.Unmarshal(bodyBytes, &basic)
		dataJson, _ := json.Marshal(basic.Data)
		dataJsonStr := string(dataJson)
		timeStamp := basic.Timestamp
		privateKey := manager.TP_SERVER_CONFIG.Auth.AuthBasic.PrivateKey

		// token算法: 对称hash加密
		// token = md5(md5(bodyStr) + privateKey + timestamp)
		preStr := fmt.Sprintf("%s%s%s", utils.GetMD5Hash(dataJsonStr), privateKey,
			strconv.Itoa(timeStamp))
		genToken := utils.GetMD5Hash(preStr)
		if genToken != basic.Token {
			response.FailWithMessage("token error, this is a illegal action.", c)
			manager.TP_LOG.Error("token error, this is a illegal action.",
				zap.String("data", string(bodyBytes)),
				zap.String("req-token", basic.Token),
				zap.String("gen-token", genToken))

			c.Abort()
			return
		}
		c.Next()
		fmt.Println("next auth basic...")
	}
}
