package model

import "time"

// User 数据库user结构体
type User struct {
	Id           int64     `gorm:"column:id"`
	Username     string    `gorm:"column:username"`
	UserAccount  string    `gorm:"column:userAccount"`
	AvatarUrl    string    `gorm:"column:avatarUrl"`
	Gender       int8      `gorm:"column:gender"`
	UserPassword string    `gorm:"column:userPassword"`
	Phone        string    `gorm:"column:phone"`
	Email        string    `gorm:"column:email"`
	UserStatus   int       `gorm:"column:userStatus"`
	UserRole     int       `gorm:"column:userRole"`
	CreateTime   time.Time `gorm:"column:createTime;autoCreateTime"`
	UpdateTime   time.Time `gorm:"column:updateTime;autoUpdateTime"`
	IsDelete     int8      `gorm:"column:isDelete"`
}
