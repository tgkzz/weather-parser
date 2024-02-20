package handler

import "github.com/gin-gonic/gin"

func (h *Handler) Routes() *gin.Engine {
	r := gin.New()

	r.GET("/weather", h.getCity)
	r.PUT("/weather", h.putData)

	return r
}
