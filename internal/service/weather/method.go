package weather

import (
	"database/sql"
	"encoding/json"
	"errors"
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
		return err
	}
	defer resp.Body.Close()

	// change error to another instance
	if resp.StatusCode != http.StatusOK {
		return errors.New("bad status code")
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
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			if err = w.repo.CreateNewData(city); err != nil {
				return err
			}
		} else {
			return err
		}
	} else {
		if err = w.repo.UpdateCityByModel(city); err != nil {
			return err
		}
	}

	return nil
}
