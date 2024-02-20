package weather

import "weather/internal/model"

// CREATE
func (w *WeatherRepo) CreateNewData(city model.City) error {
	query := `INSERT INTO weather (city, temp, tempFahrenheit, tempKelvin, main, description, last_updated)
	VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err := w.DB.Exec(query, city.CityName, city.Temp, city.TempFahrenheit, city.TempKelvin, city.Main, city.Description, city.LastUpdated)
	return err
}

// READ
func (w *WeatherRepo) GetCityByName(cityName string) (model.City, error) {
	var city model.City

	query := `SELECT id, city, temp, tempFahrenheit, tempKelvin, main, description, last_updated FROM weather WHERE city = $1`

	row := w.DB.QueryRow(query, cityName)
	if err := row.Scan(&city.Id, &city.CityName, &city.Temp, &city.TempFahrenheit, &city.TempKelvin, &city.Main, &city.Description, &city.LastUpdated); err != nil {
		return model.City{}, err
	}

	return city, nil
}

// UPDATE
func (w *WeatherRepo) UpdateCityByModel(city model.City) error {
	query := `UPDATE weather SET temp = $2, tempFahrenheit = $3, tempKelvin = $4, main = $5, description = $6, last_updated = $7 WHERE city = $1`
	_, err := w.DB.Exec(query, city.CityName, city.Temp, city.TempFahrenheit, city.TempKelvin, city.Main, city.Description, city.LastUpdated)
	return err
}

// DELETE
func (w *WeatherRepo) DeleteCityByCityName(cityName string) error {
	query := `DELETE FROM weather WHERE city = $1`
	_, err := w.DB.Exec(query, cityName)
	return err
}
