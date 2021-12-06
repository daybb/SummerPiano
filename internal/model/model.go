package model

import (
	"gorm.io/gorm"
)

type BaseModel struct {
	Id        uint32 `gorm:"primary_key"`
	CreatedAt int64
	UpdatedAt int64
	DeletedAt gorm.DeletedAt `sql:"index"`
}
