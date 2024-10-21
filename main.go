package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/IONOS-Forecast/gocast-development-ahmad/api/server"
	"github.com/IONOS-Forecast/gocast-development-ahmad/pkg/api"
	"github.com/IONOS-Forecast/gocast-development-ahmad/pkg/db"
	"github.com/IONOS-Forecast/gocast-development-ahmad/pkg/model"
	"github.com/IONOS-Forecast/gocast-development-ahmad/pkg/output"
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

	citynumbers := api.GetCityData(city, CityAPIKey)
	date := CreateDate(year, month, day)

	var weather_records model.WeatherDataForDay
	weather_records = db.ReceiveWeatherDataFromDB(date, city)

	if len(weather_records.WeatherDataForTheDay) == 0 {
		weather_records = api.GetWeatherDataFromAPI(date, citynumbers)
		db.InsertDataToDB(weather_records, citynumbers)
		output.SaveWeatherDataAsJSON(date, weather_records, citynumbers)
	}

	output.PrintWeather(weather_records, hour)
	//
	reg := prometheus.NewRegistry() // does nothing
	server.FirstMetric(reg)         //register global variables in the registry above

	go func() {
		c := time.NewTicker(1 * time.Hour)
		for {
			select {
			case <-c.C:
				now := time.Now().Format("2006-01-02")
				weather_records = api.GetWeatherDataFromAPI(now, citynumbers)
				server.SetValueFirstMetricNow(weather_records)
			}
		}
	}()

	promHandler := promhttp.HandlerFor(reg, promhttp.HandlerOpts{}) // does nothing, expose the global variables humudity etc..
	pMux := http.NewServeMux()
	pMux.Handle("/metrics", promHandler)
	//

	//http.HandleFunc("/metrics", promhttp.Handler().ServeHTTP)
	go func() {
		http.ListenAndServe(":8081", pMux)
	}()

	go func() {
		http.HandleFunc("/", h.Handler)
		http.ListenAndServe(":8080", nil)
	}()

	select {}
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
