package base

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//GetURL 透過HTTP GET取得網頁資料
func GetURL(url string) []byte {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	client := &http.Client{}
	result, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	body, _ := ioutil.ReadAll(result.Body)
	defer result.Body.Close()
	return body
}

//JSONDecode 將網頁回傳之文字資料轉成JSON格式
func JSONDecode(body []byte) interface{} {
	var resultObject interface{}
	json.Unmarshal(body, &resultObject)
	return resultObject
}
