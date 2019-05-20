package main

import (
	"fmt"

	db "zipcode/database"
	hook "zipcode/hooks"
	model "zipcode/models"
)

func init() {
	if db.Db.HasTable(&model.ZipCode{}) {
		db.Db.DropTable(&model.ZipCode{})
	}
	db.Db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.ZipCode{})
}

func main() {
	channel := make(chan int)
	go func() { channel <- hook.UpdateZipCode() }()
	result := <-channel
	fmt.Println(result)
}
