package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"hiris/models"
)

func main() {
	db, err := gorm.Open("mysql", "root:123456@(192.168.100.133)/hiris?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {

	}
	//创建users表
	//user := models.User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}
	//db.CreateTable(&models.User{})
	//db.NewRecord(user)

	//db.Create(&user)
	//fmt.Println(time.Now().Unix())

	user := models.User{}
	db.First(&user)

	fmt.Println(user.Birthday)

	defer db.Close()
}
