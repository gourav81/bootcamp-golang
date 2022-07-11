//main.go
package main

import (
	"E3/Config"
	"E3/Models"
	"E3/Routes"
	"fmt"
	"github.com/jinzhu/gorm"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Order{}, &Models.Product{})
	r := Routes.SetupRouter()
	//running
	r.Run()
}
