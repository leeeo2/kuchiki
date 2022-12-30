package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ModelTime
	ControlBy
	UserId   int    `gorm:"primaryKey;autoIncrement;comment:用户ID"  json:"userId"`
	Username string `json:"username" gorm:"size:64;comment:用户名"`
	Password string `json:"-" gorm:"size:128;comment:密码"`
	NickName string `json:"nickName" gorm:"size:128;comment:昵称"`
	Phone    string `json:"phone" gorm:"size:11;comment:手机号"`
	RoleId   int    `json:"roleId" gorm:"size:20;comment:角色ID"`
	Salt     string `json:"-" gorm:"size:255;comment:加盐"`
	Avatar   string `json:"avatar" gorm:"size:255;comment:头像"`
	Sex      string `json:"sex" gorm:"size:255;comment:性别"`
	Email    string `json:"email" gorm:"size:128;comment:邮箱"`
	Remark   string `json:"remark" gorm:"size:255;comment:备注"`
	Status   string `json:"status" gorm:"size:4;comment:状态"`
}

func (u *User) TableName() string {
	return "user"
}

func (u *User) EncryptPasswd() (err error) {
	if u.Password == "" {
		return
	}
	var hash []byte
	if hash, err = bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost); err != nil {
		return
	} else {
		u.Password = string(hash)
		return
	}
}

func (u *User) BeforeCreate(_ *gorm.DB) error {
	return u.EncryptPasswd()
}

func (u *User) BeforeUpdate(_ *gorm.DB) error {
	if u.Password != "" {
		return u.EncryptPasswd()
	}
	return nil
}
