package hooks

import (
	base "../base"
	env "../env"
	model "../models"
)

//UpdateZipCode 連上中華郵政處理回傳之行政區域資料寫入資料表
func UpdateZipCode() int {
	var cfg = env.GetEnv()
	url := cfg.Hooks.ZipCodeURL
	body := base.GetURL(url)
	resultObject := base.JSONDecode(body)
	items := resultObject.([]interface{})
	for _, value := range items {
		code := value.(map[string]interface{})["_x0033_碼郵遞區號"].(string)
		name := []rune(value.(map[string]interface{})["行政區名"].(string))
		city := string(name[:3])
		area := string(name[3:])
		list := model.ZipCode{Code: code, City: city, Area: area}
		model.AddZipCode(&list)
	}
	return 1
}
