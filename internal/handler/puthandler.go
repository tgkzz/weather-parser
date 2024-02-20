package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"weather/internal/model"
)

func (h *Handler) putData(c *gin.Context) {
	cityName := c.Query("city")
	if cityName == "" {
		ErrorHandler(c, model.ErrEmptyParams, http.StatusBadRequest)
		return
	}

	//handle different errors
	if err := h.service.Weather.InsertData(cityName); err != nil {
		if errors.Is(err, model.ErrNoCity) {
			ErrorHandler(c, err, http.StatusNotFound)
		} else {
			ErrorHandler(c, err, http.StatusInternalServerError)
		}
		return
	}

	SuccessHandler(c, model.SuccessPutOperation)
}
