package weather

import (
	"database/sql"
	"weather/internal/model"
)

type WeatherRepo struct {
	DB *sql.DB
}

type IWeatherRepo interface {
	CreateNewData(city model.City) error
	GetCityByName(cityName string) (model.City, error)
	GetAllCities() ([]string, error)
	UpdateCityByModel(city model.City) error
	DeleteCityByCityName(cityName string) error
}

func NewWeatherRepo(db *sql.DB) *WeatherRepo {
	return &WeatherRepo{DB: db}
}
