package server

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/IONOS-Forecast/gocast-development-ahmad/pkg/model"
)

const link = "http://127.0.0.1:8080/"

var dates_server = []string{

	"2024-09-01",
	"2024-09-02",
	"2024-09-03",
	"2024-09-04",
	"2024-09-05",
	"2024-09-06",
	"2024-09-07",
}

var citiesMap = map[string]string{

	"berlin":  "",
	"hamburg": "",
	"munich":  "",
}

var citiesSlice = []string{
	"berlin",
	"hamburg",
	"munich",
}

var datesInRange = map[string]string{
	"2024-09-01": "",
	"2024-09-02": "",
	"2024-09-03": "",
	"2024-09-04": "",
	"2024-09-05": "",
	"2024-09-06": "",
	"2024-09-07": "",
}

func readDataFromFile(date string, city string) model.WeatherDataForDay {

	var weatherrecords model.WeatherDataForDay
	day, month, year := DateParse(date)

	if year != 2024 {
		return weatherrecords
	}

	filename := fmt.Sprintf("%s_%.2d-%.2d-orig.json", strings.ToLower(city), day, month)

	file, err := os.ReadFile("../../resources/weather_records/testing/" + filename)
	if err != nil {
		return weatherrecords
	}

	err = json.Unmarshal(file, &weatherrecords)
	if err != nil {
		ErrorPrinting(err)
	}

	return weatherrecords
}

func DateParse(dateS string) (day int, month int, year int) {
	date, err := time.Parse("2006-01-02", dateS)

	if err != nil {
		ErrorPrinting(err)
	}

	return date.Day(), int(date.Month()), date.Year()
}

func TestOKWithData(t *testing.T) {

	go startServer()

	for i := 0; i < 50; i++ {
		t.Run(strconv.Itoa(i+1), func(t *testing.T) {

			tempDate := dates_server[rand.Intn(7)]
			tempCity := citiesSlice[rand.Intn(3)]
			var tempWeather_records model.WeatherDataForDay

			tempURL, err := url.Parse(link)
			ErrorPrinting(err)

			tempValues := tempURL.Query()
			tempValues.Set("date", tempDate)
			tempValues.Set("city", tempCity)

			tempURL.RawQuery = tempValues.Encode()

			resp, err := http.Get(tempURL.String())
			ErrorPrinting(err)

			body, err := io.ReadAll(resp.Body)
			ErrorPrinting(err)

			_ = json.Unmarshal(body, &tempWeather_records)

			if resp.StatusCode != 200 {
				t.Errorf("Test failed\nexpected response code:200\noutput response code:%d", resp.StatusCode)
			} else if len(tempWeather_records.WeatherDataForTheDay) != 25 { //25 bacause the API takes an extra hour from the next day (00:00) and adds it to the reponse
				t.Errorf("Test failed\narray weather_records isnt full. Array elements:%d", len(tempWeather_records.WeatherDataForTheDay))
			}
		})
	}
}

func TestOkWithoutData(t *testing.T) {

	//go startServer()

	for i := 0; i < 50; i++ {
		t.Run(strconv.Itoa(i+1), func(t *testing.T) {

			tempDate := randomDate()
			tempCity := citiesSlice[rand.Intn(3)]
			var tempWeather_records model.WeatherDataForDay

			tempURL, err := url.Parse(link)
			ErrorPrinting(err)

			tempValues := tempURL.Query()
			tempValues.Set("date", tempDate)
			tempValues.Set("city", tempCity)

			tempURL.RawQuery = tempValues.Encode()

			resp, err := http.Get(tempURL.String())
			ErrorPrinting(err)

			body, err := io.ReadAll(resp.Body)
			ErrorPrinting(err)

			_ = json.Unmarshal(body, &tempWeather_records)

			if resp.StatusCode != 200 {
				t.Errorf("Test failed\nexpected response code:200\noutput response code:%d", resp.StatusCode)
			} else if len(tempWeather_records.WeatherDataForTheDay) != 0 { //25 bacause the API takes an extra hour from the next day (00:00) and adds it to the reponse
				t.Errorf("Test failed\narray weather_records expected to be empty. Array elements:%d\ncity:%s date:%s", len(tempWeather_records.WeatherDataForTheDay), tempCity, tempDate)
			}
		})
	}
}

func TestNotFound(t *testing.T) {

	//go startServer()

	for i := 0; i < 50; i++ {
		t.Run(strconv.Itoa(i+1), func(t *testing.T) {
			tempDate := randomDate()
			tempCity := cityOutOfRange()

			var tempWeather_records model.WeatherDataForDay

			tempURL, err := url.Parse(link)
			ErrorPrinting(err)

			tempValues := tempURL.Query()
			tempValues.Set("date", tempDate)
			tempValues.Set("city", tempCity)

			tempURL.RawQuery = tempValues.Encode()

			resp, err := http.Get(tempURL.String())
			ErrorPrinting(err)

			body, err := io.ReadAll(resp.Body)
			ErrorPrinting(err)

			_ = json.Unmarshal(body, &tempWeather_records)

			if resp.StatusCode != 404 {
				t.Errorf("Test failed\nexpected response code:404\noutput response code:%d", resp.StatusCode)
			}
		})
	}

}

func startServer() {
	http.HandleFunc("/", HandlerMock)
	http.ListenAndServe(":8080", nil)
	/* select {} */
}

func HandlerMock(w http.ResponseWriter, req *http.Request) {

	date := req.URL.Query().Get("date")
	city := strings.ToLower(req.URL.Query().Get("city"))

	_, err := time.Parse(time.DateOnly, date)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if _, exists := citiesMap[city]; !exists {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	weather_records := readDataFromFile(date, city)

	if len(weather_records.WeatherDataForTheDay) == 0 {
		w.WriteHeader(http.StatusOK)
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

		_, err = w.Write(output)
		if err != nil {
			log.Printf("failed to write response: %+v", err)
		}
	}
}

func randomDate() string {

	min := time.Date(2010, 1, 1, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Now().Unix()
	delta := max - min

	for {
		sec := rand.Int63n(delta) + min
		date := time.Unix(sec, 0).Format(time.DateOnly)

		if _, exists := datesInRange[date]; !exists {
			return date
		}
	}
}

func cityOutOfRange() string {

	for {
		city := cities[rand.Intn(100)]
		if _, exist := citiesMap[city]; !exist {
			return city
		}
	}
}
