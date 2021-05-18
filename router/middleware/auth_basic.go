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
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
	"io/ioutil"
	"strconv"
)

func AuthBasic() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 重新写回Request Body, 供controller使用
		bodyBytes, _ := ioutil.ReadAll(c.Request.Body)
		c.Request.Body.Close()  //  must close
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

		var header request.Header
		err := mapstructure.Decode(c.Request.Header, &header)
		if err != nil {
			response.FailWithMessage("Request header data invalid.", c)
			manager.TP_LOG.Error("Request header data invalid.")
			c.Abort()
			return
		}

		var params request.Basic
		_ = json.Unmarshal(bodyBytes, &params)

		err = manager.TP_VALIDATE.Struct(params)
		if err != nil {
			response.FailWithMessage("Request data invalid.", c)
			manager.TP_LOG.Error("Request data invalid.")
			c.Abort()
			return
		}

		privateKey := manager.TP_SERVER_CONFIG.Auth.AuthBasic.PrivateKey

		dataJson, _ := json.Marshal(params.Data)
		dataJsonStr := string(dataJson)
		// token算法: 对称hash加密
		// token = md5(md5(dataJsonStr) + privateKey + timestamp)
		preStr := fmt.Sprintf("%s%s%s", utils.GetMD5Hash(dataJsonStr), privateKey,
			strconv.Itoa(params.Timestamp))
		genToken := utils.GetMD5Hash(preStr)
		
		if genToken != params.Token {
			response.FailWithMessage("token error, this is a illegal action.", c)
			manager.TP_LOG.Error("token error, this is a illegal action.",
				zap.String("data", dataJsonStr),
				zap.String("req-token", params.Token),
				zap.String("gen-token", genToken))

			c.Abort()
			return
		}
		c.Next()
	}
}
