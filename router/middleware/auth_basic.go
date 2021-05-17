package middleware

import (
	"com.pippishen/trans-proxy/manager"
	"com.pippishen/trans-proxy/model/request"
	"com.pippishen/trans-proxy/model/response"
	"com.pippishen/trans-proxy/utils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

func AuthBasic() gin.HandlerFunc {
	return func(c *gin.Context) {
		var params request.Basic
		_ = c.ShouldBindJSON(&params)
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
