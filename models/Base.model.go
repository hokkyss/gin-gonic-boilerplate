package models

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	Id        uint `gorm:"column:id;type:int;primaryKey;autoIncrement;index:,sort:asc,type:btree"`
	CreatedAt time.Time `gorm:"column:created_at;type:TIMESTAMP WITH TIME ZONE;index:,sort:asc,type:btree;default:CURRENT_TIMESTAMP;not null;"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:TIMESTAMP WITH TIME ZONE;default:CURRENT_TIMESTAMP;not null;"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index;default:null"`
}
