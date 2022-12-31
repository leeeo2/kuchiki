package models

import (
	"time"

	"github.com/leeexeo/kuchiki/pkg/common/setting"

	"github.com/leeexeo/kon/orm"
	"gorm.io/gorm"
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

var tables = []interface{}{
	&User{},
}

func Setup() error {
	c := setting.GlobalConfig()
	return orm.SetupGlobal(&c.Database, &c.Log, tables...)
}

func GetDb() *gorm.DB {
	return orm.GetDb()
}
