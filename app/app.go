package app

import (
	"fmt"

	"github.com/batuhannoz/paribu-case/app/handler"
	"github.com/batuhannoz/paribu-case/app/store/model"
	"github.com/batuhannoz/paribu-case/config"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type App struct {
	fiber *fiber.App
	db    *gorm.DB
}

func (a *App) Initialize(config *config.Config) {
	var err error
	a.db, err = gorm.Open(sqlite.Open("./database/weather.db"), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	a.db.AutoMigrate(&model.Weather{})
	a.fiber = fiber.New()
	a.registerRoutes()
}

func (a *App) registerRoutes() {
	a.fiber.Get("/weather", handler.Weather)
}

func (a *App) Run(host string) {
	a.fiber.Listen(host)
}
