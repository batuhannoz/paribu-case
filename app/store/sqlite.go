package store

import (
	"github.com/batuhannoz/paribu-case/app/store/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewSQLite() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("./weather.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&model.Weather{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
