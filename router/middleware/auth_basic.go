package middleware

import (
	"TransProxy/manager"
	"TransProxy/model/request"
	"TransProxy/model/response"
	"TransProxy/utils"
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
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

		var header request.Header
		err := mapstructure.Decode(c.Request.Header, &header)
		if err != nil {
			response.FailWithMessage("Request header data invalid.", c)
			manager.TP_LOG.Error("Request header data invalid.")
			c.Abort()
			return
		}

		bodyStr := string(bodyBytes)
		timeStamp, _ := strconv.Atoi(header.Timestamp[0])
		privateKey := manager.TP_SERVER_CONFIG.Auth.AuthBasic.PrivateKey
		// token算法: 对称hash加密
		// token = md5(md5(bodyStr) + privateKey + timestamp)
		preStr := fmt.Sprintf("%s%s%s", utils.GetMD5Hash(bodyStr), privateKey,
			strconv.Itoa(timeStamp))
		genToken := utils.GetMD5Hash(preStr)

		if genToken != header.Token[0] {
			response.FailWithMessage("token error, this is a illegal action.", c)
			manager.TP_LOG.Error("token error, this is a illegal action.",
				zap.String("data", bodyStr),
				zap.String("req-token", header.Token[0]),
				zap.String("gen-token", genToken))

			c.Abort()
			return
		}
		c.Next()
	}
}
