package handler

type WeatherResponse struct {
	Location    string  `json:"location"`
	Temperature float32 `json:"Temperature"`
}
