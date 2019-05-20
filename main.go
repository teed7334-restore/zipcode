package main

import (
	"fmt"

	db "github.com/teed7334-restore/zipcode/database"
	hook "github.com/teed7334-restore/zipcode/hooks"
	model "github.com/teed7334-restore/zipcode/models"
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
