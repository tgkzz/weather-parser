package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"weather/internal/model"
)

func (h *Handler) PutData(c *gin.Context) {
	cityName := c.Query("city")
	if cityName == "" {
		ErrorHandler(c, model.ErrEmptyParams, http.StatusBadRequest)
		return
	}

	//handle different errors
	if err := h.service.Weather.InsertData(cityName); err != nil {
		ErrorHandler(c, err, http.StatusInternalServerError)
		return
	}

	SuccessHandler(c, model.SuccessPutOperation)
}
