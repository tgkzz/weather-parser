package handler

import "weather/logger"

func (h *Handler) Scheduler(logger logger.Logger) {
	cities, err := h.service.Weather.GetAllCities()
	if err != nil {
		logger.ErrLog.Print(err)
		return
	}

	for _, city := range cities {
		if err := h.service.Weather.InsertData(city); err != nil {
			logger.ErrLog.Print(err)
		}
	}

	logger.InfoLog.Print("Worker finished its work")
}
