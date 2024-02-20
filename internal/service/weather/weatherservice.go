package weather

import "weather/internal/repository/weather"

type WeatherService struct {
	repo          weather.IWeatherRepo
	weatherApiKey string
	weatherURL    string
}

type IWeatherService interface {
	InsertData(city string) error
}

func NewWeatherService(repo weather.IWeatherRepo, weatherURL, apiKey string) *WeatherService {
	return &WeatherService{
		repo:          repo,
		weatherURL:    weatherURL,
		weatherApiKey: apiKey,
	}
}
