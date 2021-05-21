package utils

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func DoRequest(data map[string]interface{}, url string, method string) ([]byte, error) {
	bytesData, err := json.Marshal(data)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	log.Println(string(bytesData))

	reader := bytes.NewReader(bytesData)
	request, err := http.NewRequest(method, url, reader)
	defer request.Body.Close()    //程序在使用完回复后必须关闭回复的主体
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	//必须设定该参数,POST参数才能正常提交，意思是以json串提交数据

	client := http.Client{}
	resp, err := client.Do(request) //Do 方法发送请求，返回 HTTP 回复
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return respBytes, nil
}
