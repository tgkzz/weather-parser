package repository

import (
	"database/sql"
	"weather/internal/repository/weather"
)

type Repository struct {
	weather.IWeatherRepo
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		IWeatherRepo: weather.NewWeatherRepo(db),
	}
}
