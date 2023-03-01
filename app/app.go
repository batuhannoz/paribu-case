package app

import (
	"fmt"

	"github.com/batuhannoz/paribu-case/app/handler"
	"github.com/batuhannoz/paribu-case/app/service"
	"github.com/batuhannoz/paribu-case/app/store"
	"github.com/batuhannoz/paribu-case/config"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type App struct {
	fiber *fiber.App
	db    *gorm.DB
}

func (a *App) Initialize(config *config.Config) {
	var err error
	a.db, err = store.NewSQLite()
	if err != nil {
		fmt.Println(err)
	}

	weatherStore := store.NewWeatherStore(a.db)

	singleflight := service.NewSingleflight(weatherStore)

	Apphandler := handler.NewHandler(singleflight)

	a.fiber = fiber.New()
	a.registerRoutes(Apphandler)
}

func (a *App) registerRoutes(handler *handler.Handler) {
	a.fiber.Get("/weather", handler.Weather)
}

func (a *App) Run(host string) {
	a.fiber.Listen(host)
}
