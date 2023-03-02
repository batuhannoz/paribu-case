package service

import (
	"fmt"
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
	wg          sync.WaitGroup
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
				Location:  location,
				Tempature: (weather.Api1(location) + weather.Api2(location)) / 2,
			}
			s.dbChan <- &model.Weather{
				Id:         0,
				Location:   group.response.Location,
				Tempature:  group.response.Tempature,
				CreateDate: time.Now(),
			}
			fmt.Println(group.subscribers)
			group.wg.Add(int(group.subscribers))
			close(group.wait)
			group.wg.Wait()

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
	group.wg.Done()
	return &res
}
