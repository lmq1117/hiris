package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

//user表
type User struct {
	gorm.Model

	Birthday time.Time
	Age      int
	Name     string `gorm:"size:255,default:''"`
	Num      int
}

//创建之前设置字段的值
func (u *User) BeforeCreate(s *gorm.Scope) {
	s.SetColumn("Num", 123)
}
