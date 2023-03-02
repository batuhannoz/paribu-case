package model

import (
	"gorm.io/gorm"
	"time"
)

func (Weather) TableName() string {
	return "weather"
}

type Weather struct {
	gorm.Model
	ID          int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Location    string    `gorm:"column:location" json:"location"`
	Temperature float32   `gorm:"column:temperature" json:"temperature"`
	CreateDate  time.Time `gorm:"column:create_date" json:"create_date"`
}
