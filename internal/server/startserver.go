package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"weather/logger"
)

func StartServer(host, port string, r *gin.Engine, appLogger logger.Logger) error {
	srv := http.Server{
		Addr:         host + ":" + port,
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	appLogger.InfoLog.Print("Server has been started")
	if err := srv.ListenAndServe(); err != nil {
		appLogger.ErrLog.Print(err)
	}

	return nil
}
