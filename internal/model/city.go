package model

import "time"

type City struct {
	Id             int       `json:"id"`
	CityName       string    `json:"name"`
	Temp           float64   `json:"temp"`
	TempFahrenheit float64   `json:"tempFahrenheit"`
	TempKelvin     float64   `json:"tempKelvin"`
	Main           string    `json:"main"`
	Description    string    `json:"description"`
	LastUpdated    time.Time `json:"lastUpdated"`
}
