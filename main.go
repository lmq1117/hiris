package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"hiris/app"
	"hiris/app/http/models"
)

var users []models.User
var user models.User

func init() {
	initDB() //初始化db
}
func main() {
	app.InitIris()
}

func initDB() {
	//连接数据库
	models.ConnectDB()
}
