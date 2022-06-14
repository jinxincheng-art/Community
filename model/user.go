package model

import "time"

type User struct {
	Id int `gorm:"column:id"`
	//用户名
	Username string `gorm:"column:username"`
	//密码
	Password string `gorm:"column:password"`
	//加密盐
	Salt string `gorm:"column:salt"`
	//邮箱
	Email string `gorm:"column:email"`
	//类型
	Type int `gorm:"column:type"`
	//状态
	Status int `gorm:"column:status"`
	//激活码
	ActivationCode string `gorm:"column:activationCode"`
	//头地址
	HeaderUrl string `gorm:"column:headerUrl"`
	//创建时间
	CreateTime time.Time `gorm:"column:createTime"`
}

func NewUser() User {
	return User{}
}
