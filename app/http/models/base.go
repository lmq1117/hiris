package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
)

var db *gorm.DB
var err error

//打开数据库
func ConnectDB() *gorm.DB {
	if db != nil {
		return db
	}
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "" + defaultTableName
	}

	db, err = gorm.Open("mysql", "root:123lmqde@tcp(47.52.22.55:3306)/i_blog?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Errorf("open db failed %s", err.Error())
	}
	db.SingularTable(true) //全局禁用表名复数形式 true : User --> user

	db.DB().SetMaxIdleConns(5)
	db.DB().SetMaxOpenConns(20)

	db.LogMode(true)

	//db.AutoMigrate(&models.User{})

	return db
}

//关闭数据库
func CloseDB() {
	if db == nil {
		return
	}
	if err := db.Close(); err != nil {
		log.Errorf("Disconnect from database failed :%s", err.Error())
	}
}
