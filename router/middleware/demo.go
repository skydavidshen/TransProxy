package middleware

import (
	"TransProxy/model/request"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func Demo() gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyBytes, _ := ioutil.ReadAll(c.Request.Body)
		c.Request.Body.Close()  //  must close
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

		var params request.Basic
		_ = json.Unmarshal(bodyBytes, &params)
		fmt.Println("before", params.Data)

		c.Next()
		fmt.Println("after middle demo....")
	}
}
