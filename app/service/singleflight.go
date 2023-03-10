package service

import (
	"sync"
	"time"

	"github.com/batuhannoz/paribu-case/app/handler"
	"github.com/batuhannoz/paribu-case/app/store/model"
	"github.com/batuhannoz/paribu-case/app/weather"
)

type WeatherStore interface {
	SaveWeather(weather *model.Weather)
}

type Group struct {
	response    *handler.WeatherResponse
	wait        chan struct{}
	groupFull   chan struct{}
	subscribers int32
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
		m:            make(map[string]*Group, 100),
		dbChan:       make(chan *model.Weather, 1),
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
			wait:        make(chan struct{}),
			groupFull:   make(chan struct{}),
			subscribers: 0,
		}
		s.m[location] = group

		go func() {
			select {
			case <-time.After(time.Second * 5):
			case <-group.groupFull:
			}
			group.response = &handler.WeatherResponse{
				Location:    location,
				Temperature: (weather.Api1(location) + weather.Api2(location)) / 2,
			}
			s.dbChan <- &model.Weather{
				ID:          0,
				Location:    group.response.Location,
				Temperature: group.response.Temperature,
				CreateDate:  time.Now(),
			}
			close(group.wait)
		}()
	}
	group.subscribers++
	if group.subscribers >= 10 {
		delete(s.m, location)
		group.groupFull <- struct{}{}
	}
	s.mtx.Unlock()

	<-group.wait

	res := *group.response
	return &res
}
