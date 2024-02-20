package main

import (
	"log"
	"os"
	"weather/config"
	"weather/internal/handler"
	"weather/internal/repository"
	"weather/internal/server"
	"weather/internal/service"
)

func main() {
	var cfgPath string
	switch len(os.Args[1:]) {
	case 1:
		cfgPath = os.Args[1]
	case 0:
		cfgPath = "./.env"
	default:
		log.Print("USAGE: go run [CONFIG_PATH]")
		return
	}

	cfg, err := config.LoadConfig(cfgPath)
	if err != nil {
		log.Print(err)
		return
	}

	db, err := repository.LoadDB(cfg.DB)
	if err != nil {
		log.Print(err)
		return
	}

	r := repository.NewRepository(db)

	s := service.NewService(*r, cfg.WeatherApi.URL, cfg.WeatherApi.SecretKey)

	h := handler.NewHandler(s)

	if err := server.StartServer(cfg.Host, cfg.Port, h.Routes()); err != nil {
		log.Print(err)
		return
	}
}
