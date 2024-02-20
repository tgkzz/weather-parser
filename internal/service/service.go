package service

import (
	"weather/internal/repository"
	"weather/internal/service/weather"
	"weather/logger"
)

type Service struct {
	Weather weather.IWeatherService
}

func NewService(repo repository.Repository, weatherURL string, weatherApikey string, servLog logger.Logger) *Service {
	return &Service{
		Weather: weather.NewWeatherService(repo, weatherURL, weatherApikey, servLog),
	}
}
