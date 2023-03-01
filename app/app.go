package app

import (
	"github.com/batuhannoz/paribu-case/app/handler"
	"github.com/batuhannoz/paribu-case/config"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type App struct {
	fiber *fiber.App
	db *gorm.DB
}

func (a *App) Initialize(config *config.Config) {
	// TODO open database conn
	a.fiber = fiber.New()
	a.registerRoutes()
}

func (a *App) registerRoutes() {
	a.fiber.Get("/weather", handler.Weather)
}

func (a *App) Run(host string) {

}