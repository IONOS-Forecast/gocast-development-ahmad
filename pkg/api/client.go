package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/IONOS-Forecast/gocast-development-ahmad/pkg/model"
)

const CityAPIURL string = "http://api.openweathermap.org/geo/1.0/direct"
const WeatherAPIURL string = "https://api.brightsky.dev/weather"

// a function that takes city name, API key as input, gets latitude, longitude data using API and gives a struct with city name, city latitude, city longitude as output
func GetCityData(city string, CityAPIKey string) model.CityData {

	// start trying
	var citynumbers model.CityData
	temp := []model.CityData{}

	result, err := url.Parse(CityAPIURL)
	ErrorPrinting(err)

	values := result.Query()
	values.Set("q", city)
	values.Set("appid", CityAPIKey)
	values.Set("limit", "1")

	result.RawQuery = values.Encode()

	CityAPIURLWithParams := result.String()

	resp, err := http.Get(CityAPIURLWithParams)
	ErrorPrinting(err)

	body, err := io.ReadAll(resp.Body)
	ErrorPrinting(err)

	err = json.Unmarshal(body, &temp)
	ErrorPrinting(err)

	citynumbers.Name = strings.ToLower(temp[0].Name)
	citynumbers.Lat = temp[0].Lat
	citynumbers.Lon = temp[0].Lon

	return citynumbers

}

func GetWeatherDataFromAPI(date string, citydata model.CityData) model.WeatherDataForDay {

	var weather model.WeatherDataForDay

	result, err := url.Parse(WeatherAPIURL)
	ErrorPrinting(err)

	queries := result.Query()
	queries.Add("lat", strconv.FormatFloat(citydata.Lat, 'f', -1, 64))
	queries.Add("lon", strconv.FormatFloat(citydata.Lon, 'f', -1, 64))
	queries.Add("date", date)

	result.RawQuery = queries.Encode()

	WeatherAPIURLWithParams := result.String()

	resp, err := http.Get(WeatherAPIURLWithParams)
	ErrorPrinting(err)

	body, err := io.ReadAll(resp.Body)
	ErrorPrinting(err)

	err = json.Unmarshal(body, &weather)
	ErrorPrinting(err)

	// check if array is empty
	if len(weather.WeatherDataForTheDay) == 0 {
		log.Fatal("array is empty")
	} /*else if len(weather.WeatherDataForTheDay)-1 < hour { // check if array element exists
		log.Fatal("weather data for this hour is not to be found in the array")
	}*/

	return weather

}

func ErrorPrinting(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
