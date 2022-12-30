package models

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"leeeoxeo.github.com/kuchiki/pkg/common/log"
	"leeeoxeo.github.com/kuchiki/pkg/common/setting"
)

type ControlBy struct {
	CreateBy int `json:"createBy" gorm:"index;comment:创建者"`
	UpdateBy int `json:"updateBy" gorm:"index;comment:更新者"`
}

type ModelTime struct {
	CreatedAt time.Time      `json:"createdAt" gorm:"comment:创建时间"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"comment:最后更新时间"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index;comment:删除时间"`
}

//type BaseModel struct {
//	Id        string    `gorm:"type:varchar(255);not null;primaryKey"`
//	ModelTime
//}

var db *gorm.DB

var tables = []interface{}{
	&User{},
}

func Setup() error {
	c := setting.GlobalConfig().Database
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true", c.User, c.Password, c.Address, c.Port, c.Schema, c.Charset)
	options := fmt.Sprintf("ENGINE=%s DEFAULT CHARSET=%s COLLATE=%s", c.Engine, c.Charset, c.Collate)

	logger, err := log.NewGormLogger(setting.GlobalConfig().Log)
	if err != nil {
		return err
	}
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger,
	})
	if err != nil {
		return err
	}
	//使用Gorm的内置表属性创建表
	err = db.Set("gorm:table_options", options).AutoMigrate(
		tables...,
	)
	if err != nil {
		return err
	}

	sqlDb, _ := db.DB()
	sqlDb.SetMaxIdleConns(c.MaxIdleConn)
	sqlDb.SetMaxOpenConns(c.MaxOpenConn)
	return nil
}

func GetDb() *gorm.DB {
	return db
}
