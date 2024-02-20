package main

import (
	"github.com/robfig/cron/v3"
	"log"
	"os"
	"weather/config"
	"weather/internal/handler"
	"weather/internal/repository"
	"weather/internal/server"
	"weather/internal/service"
	"weather/logger"
)

func main() {
	appLogger, err := logger.NewLogger("app")
	if err != nil {
		log.Print("app will be created without logging\"")
		appLogger = logger.Logger{}
	}
	var cfgPath string
	switch len(os.Args[1:]) {
	case 1:
		cfgPath = os.Args[1]
	case 0:
		cfgPath = "./.env"
	default:
		appLogger.ErrLog.Print("USAGE: go run [CONFIG_PATH]")
		return
	}

	cfg, err := config.LoadConfig(cfgPath)
	if err != nil {
		appLogger.ErrLog.Print(err)
		return
	}

	db, err := repository.LoadDB(cfg.DB)
	if err != nil {
		appLogger.ErrLog.Print(err)
		return
	}

	r := repository.NewRepository(db)

	serviceLogger, err := logger.NewLogger("service")
	if err != nil {
		log.Print("app will be created without logging\"")
		serviceLogger = logger.Logger{}
	}

	s := service.NewService(*r, cfg.WeatherApi.URL, cfg.WeatherApi.SecretKey, serviceLogger)

	h := handler.NewHandler(s)

	// worker
	workerLogger, err := logger.NewLogger("worker")
	if err != nil {
		appLogger.ErrLog.Print("worker will be created without logging")
		workerLogger = logger.Logger{}
	}
	c := cron.New()
	if _, err := c.AddFunc(cfg.WorkerTime, func() {
		h.Scheduler(workerLogger)
	}); err != nil {
		workerLogger.ErrLog.Print(err)
		workerLogger.ErrLog.Print("Due to the error while creating worker, app will be started without worker")
	}
	c.Start()
	workerLogger.InfoLog.Print("Worker has started its work")

	if err := server.StartServer(cfg.Host, cfg.Port, h.Routes(), appLogger); err != nil {
		appLogger.ErrLog.Print(err)
		return
	}
}
