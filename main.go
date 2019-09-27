package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db, err := gorm.Open("mysql", "root:123456@(192.168.100.133)/hiris?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {

	}
	defer db.Close()
}
