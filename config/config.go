package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Host       string
	Port       string
	WorkerTime string
	MongoDB    MongoDB
	DB         DB
	WeatherApi WeatherApi
}

type WeatherApi struct {
	URL       string
	SecretKey string
}

type DB struct {
	DriverName     string
	DataSourceName string
}

type MongoDB struct {
	User     string
	Password string
	Host     string
	Name     string
}

func LoadConfig(path string) (Config, error) {
	if err := godotenv.Load(path); err != nil {
		return Config{}, err
	}

	DataSource := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"), os.Getenv("DB_NAME"),
	)

	cfg := Config{
		Host:       os.Getenv("HOST"),
		Port:       os.Getenv("PORT"),
		WorkerTime: os.Getenv("WORKER_TIME"),
		MongoDB: MongoDB{
			User:     os.Getenv("MONGO_USER"),
			Password: os.Getenv("MONGO_PSW"),
			Host:     os.Getenv("MONGO_HOST"),
			Name:     os.Getenv("MONGO_DB_NAME"),
		},
		DB: DB{
			DriverName:     os.Getenv("DB_DRIVER_NAME"),
			DataSourceName: DataSource,
		},
		WeatherApi: WeatherApi{
			URL:       os.Getenv("WEATHER_URL"),
			SecretKey: os.Getenv("WEATHER_SECRET_KEY"),
		},
	}

	return cfg, nil
}
