package handler

import "github.com/gin-gonic/gin"

func (h *Handler) Routes() *gin.Engine {
	r := gin.New()

	//r.GET("/weather")
	r.PUT("/weather", h.PutData)

	return r
}
