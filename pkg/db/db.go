package db

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/IONOS-Forecast/gocast-development-ahmad/pkg/model"
	"github.com/go-pg/pg/v10"
)

type DB struct {
	conn *pg.DB
}

func ConnectToDB(user, password, dbname string, addr string) DB {

	db := pg.Connect(&pg.Options{

		User:     user,
		Password: password,
		Database: dbname,
		Addr:     addr,
	})

	return DB{conn: db}
}

// a function that gets date, hour, city name as input and return a struct with weather information for that city, date and hour as output
func (db DB) ReceiveWeatherDataFromDB(date string, hour int, city string) model.WeatherDataForDay {

	var weather_records model.WeatherDataForDay

	_, err := db.conn.Query(&weather_records, fmt.Sprintf("SELECT * FROM weather_records WHERE timestamp::date='%s' and city='%s';", date, city))
	ErrorPrinting(err)
	//db.conn.Close()

	return weather_records
}

// a function that gets city name, city latitude and city longitude as a strcut and checks if this city is in the database, if not the city information (name, latitude, longitude) will be inserted into the database
func (db DB) CheckIfCityExistsInDB(citynumbers []model.CityData) {

	var cityexists bool
	_, err := db.conn.QueryOne(pg.Scan(&cityexists), "SELECT COUNT(name) AS citycounter FROM cities WHERE name=?", citynumbers[0].Name)
	ErrorPrinting(err)

	if !cityexists {
		_, err = db.conn.Exec("INSERT INTO cities (name, lat, lon) VALUES (?, ?, ?)", citynumbers[0].Name, citynumbers[0].Lat, citynumbers[0].Lon)
		ErrorPrinting(err)

		db.SaveCitiesAsJSON()
	}
}

// a function that gets cities information (name, latitude, longitude) from database and saves them as JSON in resources/weather_records/cities.json
func (db DB) SaveCitiesAsJSON() {

	type city struct {
		Name string
		Lat  float64
		Lon  float64
	}

	var cities []city

	db.conn.Model(&cities).Select()

	CitiesAsJSON, err := json.MarshalIndent(cities, "", " ")
	ErrorPrinting(err)

	err = os.MkdirAll("resources/weather_records/", os.ModePerm) //leon
	ErrorPrinting(err)                                           //leon

	file, err := os.Create("resources/weather_records/cities.json")
	ErrorPrinting(err)

	file.Write(CitiesAsJSON)
	file.Close()

}

// a function that takes a struct of weather information for 25 hours (each hour in a specific day and the first hour of the next day) and a strcut of city information (name, latitude, longitude) and inserts the 24 hour data for that specific day into the database
func (db DB) InsertDataToDB(WeatherInfo model.WeatherDataForDay, citynumbers []model.CityData) {

	var count int
	_, err := db.conn.QueryOne(pg.Scan(&count), "SELECT COUNT(id) AS count FROM weather_records;")
	ErrorPrinting(err)

	db.CheckIfCityExistsInDB(citynumbers)

	for i := 0; i < len(WeatherInfo.WeatherDataForTheDay) && i <= 23; i++ {
		count++
		WeatherInfo.WeatherDataForTheDay[i].City = strings.ToLower(citynumbers[0].Name)
		WeatherInfo.WeatherDataForTheDay[i].ID = count
		_, err = db.conn.Model(&WeatherInfo.WeatherDataForTheDay[i]).Where("timestamp=?", WeatherInfo.WeatherDataForTheDay[i].TimeStamp).Where("city=?", WeatherInfo.WeatherDataForTheDay[i].City).SelectOrInsert()
		ErrorPrinting(err)
	}

}

// a function that prints errors
func ErrorPrinting(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
