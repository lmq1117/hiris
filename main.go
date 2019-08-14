package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"hiris/app/models"
	"time"
	//"time"
)

func init() {

}
func main() {
	db, err := gorm.Open("mysql", "root:123lmqde@tcp(47.52.22.55:3306)/i_blog?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		return
	}
	db.SingularTable(true) //全局禁用表名复数形式 true : User --> user

	//db.AutoMigrate(&models.User{})

	//insert
	user := models.User{Name: "曹丕", Age: 30, Birthday: time.Now()}
	fmt.Println(db.NewRecord(user))
	db.Create(&user)
	fmt.Println(db.NewRecord(user))

	//user := models.User{}
	//db.First(&user)
	//db.Last(&user)
	//fmt.Println(user)

	//users 是 切片 类型是models.User
	//var users []models.User
	//db.Find(&users)
	//fmt.Println(users)
	//fmt.Println(users[1].CreatedAt)

	defer db.Close()
}
