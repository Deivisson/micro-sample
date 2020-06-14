package models

import (
	"time"
)

type BaseModel struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updatedAt"`
	Errors    Errors    `json:"-"`
}

type Errors struct {
	Exception      error
	Business       interface{}
	RecordNotFound error
}
