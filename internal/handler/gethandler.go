package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"weather/internal/model"
)

func (h *Handler) getCity(c *gin.Context) {
	cityName := c.Query("city")
	if cityName == "" {
		ErrorHandler(c, model.ErrEmptyParams, http.StatusBadRequest)
		return
	}

	result, err := h.service.Weather.GetCityData(cityName)
	if err != nil {
		if errors.Is(err, model.ErrNoCity) {
			ErrorHandler(c, err, http.StatusNotFound)
		} else {
			ErrorHandler(c, err, http.StatusInternalServerError)
		}
		return
	}

	c.JSON(http.StatusOK, result)
}
