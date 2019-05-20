package models

import (
	"log"

	db "zipcode/database"
)

//ZipCode 郵遞區域表
type ZipCode struct {
	ID   int `gorm:"AUTO_INCREMENT"`
	Code string
	City string
	Area string
}

//AddZipCode 新增郵遞區域
func AddZipCode(a *ZipCode) {
	err := db.Db.Create(&a).Error
	if err != nil {
		log.Fatal(err)
	}
}
