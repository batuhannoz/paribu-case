package store

import (
	"github.com/batuhannoz/paribu-case/app/store/model"
	"gorm.io/gorm"
)

type WeatherStore struct {
	connection *gorm.DB
	DBChan     chan *model.Weather
}

func NewWeatherStore(connection *gorm.DB) *WeatherStore {
	return &WeatherStore{
		connection: connection,
		DBChan:     make(chan *model.Weather, 100),
	}
}
