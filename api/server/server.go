package server

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/IONOS-Forecast/gocast-development-ahmad/pkg/db"
)

type Handler struct {
	db db.DB
}

func NewHandler(db db.DB) *Handler {
	return &Handler{db: db}
}

func (h *Handler) Handler(w http.ResponseWriter, req *http.Request) {

	date := req.URL.Query().Get("date")
	city := strings.ToLower(req.URL.Query().Get("city"))

	_, err := time.Parse(time.DateOnly, date)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("invalid timestamp format")
		return
	}

	if !h.db.CheckIfCityExistsInDB(city) {
		w.WriteHeader(http.StatusNotFound)
		log.Printf("city not found")
		return
	}

	weather_records := h.db.ReceiveWeatherDataFromDB(date, city)

	if len(weather_records.WeatherDataForTheDay) == 0 {
		w.WriteHeader(http.StatusOK)
		log.Printf("weather data for this date arent found in the database")
		return

	} else {

		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)

		output, err := json.Marshal(weather_records)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("failed to marshal response %+v\n", err)
			return
		}

		_, err = w.Write(output) //
		if err != nil {
			log.Printf("failed to write response: %+v", err)
		}
	}
}

// a function that prints errors
func ErrorPrinting(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
