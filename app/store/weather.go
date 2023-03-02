package store

import (
	"github.com/batuhannoz/paribu-case/app/store/model"
	"gorm.io/gorm"
)

type WeatherStore struct {
	connection *gorm.DB
}

func NewWeatherStore(connection *gorm.DB) *WeatherStore {
	return &WeatherStore{
		connection: connection,		
	}
}

func (w *WeatherStore) SaveWeather(weather *model.Weather) {
	w.connection.Save(&weather)
}   

