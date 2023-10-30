package models

import (
	"time"

	"github.com/leeexeo/kuchiki/pkg/common/setting"

	"github.com/leeexeo/kon/orm"
	"gorm.io/gorm"
)

type ControlBy struct {
	CreateBy string `json:"CreateBy" gorm:"size:64;index;comment:创建者ID"`
	UpdateBy string `json:"UpdateBy" gorm:"size:64;index;comment:更新者ID"`
}

type ModelTime struct {
	CreatedAt time.Time      `json:"CreatedAt" gorm:"comment:创建时间"`
	UpdatedAt time.Time      `json:"UpdatedAt" gorm:"comment:最后更新时间"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index;comment:删除时间"`
}

var tables = []interface{}{
	&User{},
	&Article{},
}

func Setup() error {
	c := setting.GlobalConfig()
	return orm.SetupGlobal(&c.Database, &c.Log, tables...)
}

func GetDb() *gorm.DB {
	return orm.GetDb()
}

type DescribeInput struct {
	PageSize   int
	PageNumber int
	Query      string
	Params     []interface{}
	Order      string
}
