package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/IONOS-Forecast/gocast-development-ahmad/api/server"
	"github.com/IONOS-Forecast/gocast-development-ahmad/pkg/api"
	"github.com/IONOS-Forecast/gocast-development-ahmad/pkg/db"
	"github.com/IONOS-Forecast/gocast-development-ahmad/pkg/model"
	"github.com/IONOS-Forecast/gocast-development-ahmad/pkg/output"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	db := db.ConnectToDB(os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_ADDRESS"))
	h := server.NewHandler(db)

	CityAPIKey := os.Getenv("CITY_API_KEY")
	city := os.Getenv("CITY")                    // city name
	year, err := strconv.Atoi(os.Getenv("YEAR")) // year
	ErrorPrinting(err)
	month, err := strconv.Atoi(os.Getenv("MONTH")) // month
	ErrorPrinting(err)
	day, err := strconv.Atoi(os.Getenv("DAY")) // day
	ErrorPrinting(err)
	hour, err := strconv.Atoi(os.Getenv("HOUR")) // hour
	ErrorPrinting(err)
	//n_minutes, err := strconv.Atoi(os.Getenv("N_MINUTES")) //n minutes
	//ErrorPrinting(err)

	//for {

	citynumbers := api.GetCityData(city, CityAPIKey)
	date := CreateDate(year, month, day)

	var weather_records model.WeatherDataForDay
	weather_records = db.ReceiveWeatherDataFromDB(date, city)

	if len(weather_records.WeatherDataForTheDay) == 0 {
		weather_records = api.GetWeatherDataFromAPI(date, hour, citynumbers)
		db.InsertDataToDB(weather_records, citynumbers)
		//output.SaveWeatherDataAsJSON() instead of body, a struct

	}

	output.PrintWeather(weather_records, hour)

	http.HandleFunc("/", h.Handler)
	http.HandleFunc("/metrics", promhttp.Handler().ServeHTTP)
	http.ListenAndServe(":8080", nil)

	//GetWeatherDataFromAPI(year, month, day, hour)

	//	time.Sleep(time.Second * time.Duration(n_minutes))

	//}
}

// a function that returns a formatted date (yyyy-mm-dd) as a string using the input year, month, day
func CreateDate(year int, month int, day int) string {

	date, err := time.Parse("2006-01-02", fmt.Sprintf("%.4d-%.2d-%.2d", year, month, day))
	ErrorPrinting(err)

	dateString := date.Format("2006-01-02")

	if year < 2010 {
		log.Fatal("input year should be 2010 or more")
	}

	return dateString

}

// a simple function for error printing
func ErrorPrinting(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
