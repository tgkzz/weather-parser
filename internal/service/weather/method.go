package weather

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net"
	"net/http"
	"net/url"
	"time"
	"weather/internal/model"
	"weather/internal/pkg"
)

func (w WeatherService) InsertData(cityName string) error {
	parsedURL, err := url.Parse(w.weatherURL)
	if err != nil {
		return err
	}

	values := parsedURL.Query()
	values.Add("q", cityName)
	values.Add("appid", w.weatherApiKey)
	values.Add("units", "metric")
	parsedURL.RawQuery = values.Encode()

	resp, err := http.Get(parsedURL.String())
	if err != nil {
		if _, ok := err.(net.Error); ok {
			return model.ErrNoInternetConnection
		}
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusTooManyRequests {
			return model.ErrToManyRequestToAnotherService
		}
		return model.ErrNoCity
	}

	var result model.WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return err
	}

	city := model.City{
		CityName:       result.Name,
		Temp:           result.Main.Temp,
		TempFahrenheit: pkg.CelsiusToFahrenheit(result.Main.Temp),
		TempKelvin:     pkg.CelsiusToKelvin(result.Main.Temp),
		Main:           result.Weather[0].Main,
		Description:    result.Weather[0].Description,
		LastUpdated:    time.Now(),
	}

	// send info to repo layer
	_, err = w.repo.GetCityByName(city.CityName)
	if errors.Is(err, sql.ErrNoRows) {
		if err := w.repo.CreateNewData(city); err != nil {
			return err
		}
		w.logger.InfoLog.Print(model.SuccessPutOperation)
		return nil
	} else if err != nil {
		return err
	}

	if err := w.repo.UpdateCityByModel(city); err != nil {
		return err
	}

	w.logger.InfoLog.Print(model.SuccessPutOperation)
	return nil
}

func (w WeatherService) GetCityData(cityName string) (model.City, error) {
	result, err := w.repo.GetCityByName(cityName)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.City{}, model.ErrNoCity
		}
		return model.City{}, err
	}

	return result, nil
}

func (w WeatherService) GetAllCities() ([]string, error) {
	res, err := w.repo.GetAllCities()
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []string{}, model.ErrNoCity
		}
		return nil, err
	}

	w.logger.InfoLog.Print(model.SuccessGetOperation)
	return res, nil
}
