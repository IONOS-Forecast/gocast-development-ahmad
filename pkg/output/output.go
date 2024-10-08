package output

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/IONOS-Forecast/gocast-development-ahmad/pkg/model"
)

// a function that takes weather information for each hour in a specific day as a struct and prints out the weather information for a specific hour
func PrintWeather(WeatherInfo model.WeatherDataForDay, hour int) {

	//fmt.Println("\033[2J")
	fmt.Printf("time:					%.16v\n", WeatherInfo.WeatherDataForTheDay[hour].TimeStamp)
	fmt.Printf("condition:				%s\n", WeatherInfo.WeatherDataForTheDay[hour].Condition)
	fmt.Printf("temperature:				%.1f\n", WeatherInfo.WeatherDataForTheDay[hour].Temperature)
	fmt.Printf("wind speed:				%.1f\n", WeatherInfo.WeatherDataForTheDay[hour].WindSpeed)
	fmt.Printf("wind direction:				%d\n", WeatherInfo.WeatherDataForTheDay[hour].WindDirection)
	fmt.Printf("wind gust speed:			%.1f\n", WeatherInfo.WeatherDataForTheDay[hour].WindGustSpeed)
	fmt.Printf("wind gust direction:			%d\n", WeatherInfo.WeatherDataForTheDay[hour].WindGustDirection)
	fmt.Printf("relative humidity:			%d\n", WeatherInfo.WeatherDataForTheDay[hour].RelativeHumidity)
	fmt.Printf("dew point:				%.1f\n", WeatherInfo.WeatherDataForTheDay[hour].DewPoint)
	fmt.Printf("precipitation probability:		%.1f\n", WeatherInfo.WeatherDataForTheDay[hour].PrecipitationProbability)
	fmt.Printf("precipitation probability 6h:		%.1f\n", WeatherInfo.WeatherDataForTheDay[hour].PrecipitationProbability6h)
	fmt.Printf("visibility:				%d\n", WeatherInfo.WeatherDataForTheDay[hour].Visibility)
	fmt.Printf("pressure in MSL:			%.1f\n", WeatherInfo.WeatherDataForTheDay[hour].PressureMsl)
	fmt.Printf("cloud cover:				%.2f\n", WeatherInfo.WeatherDataForTheDay[hour].CloudCover)
	fmt.Printf("sunshine:				%.0f\n", WeatherInfo.WeatherDataForTheDay[hour].Sunshine)
	fmt.Printf("solar:					%.3f\n", WeatherInfo.WeatherDataForTheDay[hour].Solar)
	fmt.Printf("general:				%s\n", WeatherInfo.WeatherDataForTheDay[hour].Icon)
	fmt.Printf("precipitation:				%.1f\n", WeatherInfo.WeatherDataForTheDay[hour].Precipitation)

}

// a function that takes year, month, day, body from REST API response, city information as struct and saves those weather information for the specific day as JSON in resources/weather_records/ with the city name, date and sometimes year in the filename
func SaveWeatherDataAsJSON(year int, month int, day int, body []byte, citynumbers []model.CityData) {

	var FileName string

	if year == time.Now().Year() {
		FileName = fmt.Sprintf(citynumbers[0].Name+"_%.2d-%.2d-orig.json", day, month)
	} else {
		FileName = fmt.Sprintf(citynumbers[0].Name+"_%.2d-%.2d-%d-orig.json", day, month, year)
	}

	err := os.MkdirAll("resources/weather_records/", os.ModePerm) // leon
	ErrorPrinting(err)                                            // leon

	file, err := os.Create("resources/weather_records/" + FileName)
	ErrorPrinting(err)

	file.Write(body)
	file.Close()

}

// a function that prints errors
func ErrorPrinting(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
