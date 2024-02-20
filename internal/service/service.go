package service

import (
	"weather/internal/repository"
	"weather/internal/service/weather"
)

type Service struct {
	Weather weather.IWeatherService
}

func NewService(repo repository.Repository, weatherURL string, weatherApikey string) *Service {
	return &Service{
		Weather: weather.NewWeatherService(repo, weatherURL, weatherApikey),
	}
}
