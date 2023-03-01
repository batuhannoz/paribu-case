package handler

type WeatherResponse struct {
	Location string `json:"location"`
	Tempature float32 `json:"Tempature"`
}