package server

import (
	"time"

	"github.com/IONOS-Forecast/gocast-development-ahmad/pkg/model"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	Temperature *prometheus.GaugeVec
	Humidity    *prometheus.GaugeVec
	WindSpeed   *prometheus.GaugeVec
	Pressure    *prometheus.GaugeVec
)

func FirstMetric(reg prometheus.Registerer) {
	Temperature = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "gocast",
		Name:      "temperature",
		Help:      "Temperature of each city",
	}, []string{"location", "timestamp"})
	Humidity = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "gocast",
		Name:      "humidity",
		Help:      "Humidity of each city",
	}, []string{"location", "timestamp"})
	WindSpeed = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "gocast",
		Name:      "wind_speed",
		Help:      "Wind speed of each city",
	}, []string{"location", "timestamp"})
	Pressure = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "gocast",
		Name:      "pressure",
		Help:      "Pressure of each city",
	}, []string{"location", "timestamp"})

	reg.MustRegister(Temperature)
	reg.MustRegister(Humidity)
	reg.MustRegister(WindSpeed)
	reg.MustRegister(Pressure)
}

func SetValueFirstMetric(weather model.WeatherDataForDay, hour int) {

	city := weather.WeatherDataForTheDay[hour].City
	now := weather.WeatherDataForTheDay[hour].TimeStamp.Format(time.RFC3339)
	temp := weather.WeatherDataForTheDay[hour].Temperature

	Temperature.With(prometheus.Labels{
		"location":  city,
		"timestamp": now,
	}).Set(temp)

	hum := float64(weather.WeatherDataForTheDay[hour].RelativeHumidity)
	Humidity.With(prometheus.Labels{
		"location":  city,
		"timestamp": now,
	}).Set(hum)

	speed := weather.WeatherDataForTheDay[hour].WindSpeed
	WindSpeed.With(prometheus.Labels{
		"location":  city,
		"timestamp": now,
	}).Set(speed)

	pressure := weather.WeatherDataForTheDay[hour].PressureMsl
	Pressure.With(prometheus.Labels{
		"location":  city,
		"timestamp": now,
	}).Set(pressure)
}

func SetValueFirstMetricNow(weather model.WeatherDataForDay) {
	now := time.Now()
	hour := now.Hour()

	city := weather.WeatherDataForTheDay[hour].City
	timestamp := now.Format(time.RFC3339)
	temp := weather.WeatherDataForTheDay[hour].Temperature

	Temperature.With(prometheus.Labels{
		"location":  city,
		"timestamp": timestamp,
	}).Set(temp)

	hum := float64(weather.WeatherDataForTheDay[hour].RelativeHumidity)
	Humidity.With(prometheus.Labels{
		"location":  city,
		"timestamp": timestamp,
	}).Set(hum)

	speed := weather.WeatherDataForTheDay[hour].WindSpeed
	WindSpeed.With(prometheus.Labels{
		"location":  city,
		"timestamp": timestamp,
	}).Set(speed)

	pressure := weather.WeatherDataForTheDay[hour].PressureMsl
	Pressure.With(prometheus.Labels{
		"location":  city,
		"timestamp": timestamp,
	}).Set(pressure)

}
