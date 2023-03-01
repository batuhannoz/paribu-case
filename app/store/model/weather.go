package model

import "time"

func (Weather) TableName() string {
	return "weather"
}

type Weather struct {
	Id          uint64    `gorm:"primary_key:auto_increment" json:"id"`
	City   string    `gorm:"type:varchar(50)" json:"city"`
	Tempature   float32    `gorm:"type:float32" json:"tempature"`
	CreateDate  time.Time `gorm:"type:datetime" json:"create_date"`
}