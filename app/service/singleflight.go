package service

import "github.com/batuhannoz/paribu-case/app/handler"

type WeatherStore interface {
	//
}

type Group struct {
}

type Singleflight struct {
	weatherStore WeatherStore
}

func NewSingleflight(weatherStore WeatherStore) *Singleflight {
	return &Singleflight{
		weatherStore: weatherStore,
	}
}

func (s *Singleflight) WeatherByLocation(location string) *handler.WeatherResponse {
	return &handler.WeatherResponse{}
}
