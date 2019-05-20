package hooks

import (
	"encoding/json"

	"github.com/teed7334-restore/zipcode/base"
	bean "github.com/teed7334-restore/zipcode/beans"
	"github.com/teed7334-restore/zipcode/env"
	model "github.com/teed7334-restore/zipcode/models"
)

//UpdateZipCode 連上中華郵政處理回傳之行政區域資料寫入資料表
func UpdateZipCode() int {
	var cfg = env.GetEnv()
	url := cfg.Hooks.ZipCodeURL
	body := base.GetURL(url)
	resultObject := []*bean.ZipCode{}
	json.Unmarshal(body, &resultObject)
	for _, value := range resultObject {
		code := value.GetCode()
		name := []rune(value.GetName())
		city := string(name[:3])
		area := string(name[3:])
		list := model.ZipCode{Code: code, City: city, Area: area}
		model.AddZipCode(&list)
	}
	return 1
}
