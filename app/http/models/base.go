package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
)

func ConnectDB() *gorm.DB {
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "" + defaultTableName
	}

	db, err := gorm.Open("mysql", "root:123lmqde@tcp(47.52.22.55:3306)/i_blog?charset=utf8mb4&parseTime=True&loc=Local")
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
