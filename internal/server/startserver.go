package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func StartServer(host, port string, r *gin.Engine) error {
	srv := http.Server{
		Addr:         host + ":" + port,
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	err := srv.ListenAndServe()

	return err
}
