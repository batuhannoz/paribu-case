package service

import (
	"sync"
	"time"

	"github.com/batuhannoz/paribu-case/app/handler"
	"github.com/batuhannoz/paribu-case/app/store/model"
)

type WeatherStore interface {
	SaveWeather(weather *model.Weather)
}

type Group struct {
	response     *model.Weather
	wait         chan struct{}
	groupFull    chan struct{}
	subsicribers int32
}

type Singleflight struct {
	weatherStore WeatherStore
	m            map[string]*Group
	dbChan       chan *model.Weather
	mtx          sync.Mutex
}

func NewSingleflight(weatherStore WeatherStore) *Singleflight {
	singleflight := &Singleflight{
		weatherStore: weatherStore,
		m:            make(map[string]*Group),
		dbChan:       make(chan *model.Weather, 50),
	}
	go func() {
		for {
			weather := <-singleflight.dbChan
			singleflight.weatherStore.SaveWeather(weather)
		}
	}()
	return singleflight
}

func (s *Singleflight) WeatherByLocation(location string) *handler.WeatherResponse {
	s.mtx.Lock()
	group, ok := s.m[location]
	if !ok {
		group = &Group{
			wait:         make(chan struct{}),
			groupFull:    make(chan struct{}),
			subsicribers: 0,
		}
		go func() {
			select {
			case <-time.After(time.Second * 5):
			case <-group.groupFull:
			}
			//
		}()
	}
	group.subsicribers++
	s.mtx.Unlock()
	return &handler.WeatherResponse{}
}
