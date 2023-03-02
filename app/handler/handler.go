package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Singleflight interface {
	WeatherByLocation(location string) *WeatherResponse
}

type Handler struct {
	singleflight Singleflight
}

func NewHandler(singleflight Singleflight) *Handler {
	return &Handler{
		singleflight,
	}
}

func (app *Handler) Weather(ctx *fiber.Ctx) error {
	location := ctx.Query("q")
	if location == "" {
		return ctx.SendStatus(http.StatusBadRequest)
	}
	res := app.singleflight.WeatherByLocation(location)
	return ctx.Status(http.StatusOK).JSON(res)
}
