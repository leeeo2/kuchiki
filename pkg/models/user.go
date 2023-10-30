package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	Id       string `json:"Id" gorm:"size:64;not null;primaryKey;comment:用户ID"`
	Username string `json:"Username" gorm:"size:64;comment:用户名"`
	Password string `json:"-" gorm:"size:128;comment:密码"`
	NickName string `json:"NickName" gorm:"size:128;comment:昵称"`
	Phone    string `json:"Phone" gorm:"size:11;comment:手机号"`
	RoleId   string `json:"RoleId" gorm:"size:64;comment:角色ID"`
	Salt     string `json:"-" gorm:"size:255;comment:加盐"`
	Avatar   string `json:"Avatar" gorm:"size:255;comment:头像"`
	Sex      string `json:"Sex" gorm:"size:255;comment:性别"`
	Email    string `json:"Email" gorm:"size:128;comment:邮箱"`
	Remark   string `json:"Remark" gorm:"size:255;comment:备注"`
	Status   string `json:"Status" gorm:"size:32;comment:状态"`
	ModelTime
	ControlBy
}

func (u *User) TableName() string {
	return "users"
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

type UserDao interface {
	Save(u *User) error
	Delete(id string) error
	UpdateBasic(user *User) error
	UpdateAvatar(id, avatar string) error
	UpdateStatus(id, status string) error
	DescribeUser(id string) (*User, error)
	DescribeUsers(input *DescribeInput) ([]*User, error)
}

func NewUserDao() UserDao {
	// use gorm default
	// if other
	// reuturn &OtherDaoImpl
	return &UserDaoImpl{}
}

type UserDaoImpl struct {
}

func (u *UserDaoImpl) Save(user *User) error {
	return GetDb().Create(user).Error
}

func (u *UserDaoImpl) Delete(id string) error {
	tx := GetDb()
	return tx.Where("id = ?", id).Delete(&User{}).Error
}

func (u *UserDaoImpl) UpdateBasic(user *User) error {
	tx := GetDb().Model(&User{})
	tx.Where("id = ?", user.Id)
	fields := map[string]interface{}{
		"username":  user.Username,
		"password":  user.Password,
		"nick_name": user.NickName,
		"phone":     user.Phone,
		"salt":      user.Salt,
		"sex":       user.Sex,
		"email":     user.Email,
		"remark":    user.Remark,
	}
	return tx.Updates(fields).Error
}

func (u *UserDaoImpl) UpdateAvatar(id, avatar string) error {
	tx := GetDb().Model(&User{})
	tx.Where("id = ?", id)
	update := map[string]interface{}{
		"avatar": avatar,
	}
	return tx.Updates(update).Error
}

func (u *UserDaoImpl) UpdateStatus(id, status string) error {
	tx := GetDb().Model(&User{})
	tx.Where("id = ?", id)
	update := map[string]interface{}{
		"status": status,
	}
	return tx.Updates(update).Error
}

func (u *UserDaoImpl) DescribeUser(id string) (*User, error) {
	tx := GetDb().Where("id = ?", id)

	var user User
	err := tx.First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &user, nil
}

func (u *UserDaoImpl) DescribeUsers(input *DescribeInput) ([]*User, error) {
	tx := GetDb().Where(input.Query, input.Params...)
	if input.PageSize > 0 && input.PageNumber > 0 {
		tx = tx.Limit(input.PageSize).Offset((input.PageNumber - 1) * input.PageSize)
	}

	if len(input.Order) != 0 {
		tx.Order(input.Order)
	}

	users := make([]*User, 0)
	err := tx.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}
