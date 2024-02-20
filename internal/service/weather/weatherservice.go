package weather

import (
	"weather/internal/model"
	"weather/internal/repository/weather"
	"weather/logger"
)

type WeatherService struct {
	repo          weather.IWeatherRepo
	weatherApiKey string
	weatherURL    string
	logger        logger.Logger
}

type IWeatherService interface {
	InsertData(city string) error
	GetCityData(cityName string) (model.City, error)
	GetAllCities() ([]string, error)
}

func NewWeatherService(repo weather.IWeatherRepo, weatherURL, apiKey string, logger logger.Logger) *WeatherService {
	return &WeatherService{
		repo:          repo,
		weatherURL:    weatherURL,
		weatherApiKey: apiKey,
		logger:        logger,
	}
}
