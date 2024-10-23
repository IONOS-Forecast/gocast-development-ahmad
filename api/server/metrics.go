package server

import (
	"math/rand"
	"time"

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

func UpdateMetrics() {
	city := "Berlin"
	now := time.Now().Format(time.RFC3339)
	temp := 100 * rand.Float64()
	Temperature.With(prometheus.Labels{
		"location":  city,
		"timestamp": now,
	}).Set(temp)

	hum := 140 * rand.Float64()
	Humidity.With(prometheus.Labels{
		"location":  city,
		"timestamp": now,
	}).Set(hum)

	speed := 1400 * rand.Float64()
	WindSpeed.With(prometheus.Labels{
		"location":  city,
		"timestamp": now,
	}).Set(speed)

	pressure := 2200 * rand.Float64()
	Pressure.With(prometheus.Labels{
		"location":  city,
		"timestamp": now,
	}).Set(pressure)
}

/*

func NewMetrics(reg prometheus.Registerer) *model.FirstMetricStruct {
	m = &model.FirstMetricStruct{
		Temperature: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "gocast",
			Name:      "temperature",
			Help:      "Temperature of each city",
		}, []string{"location", "timestamp"}),
		Humidity: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "gocast",
			Name:      "humidity",
			Help:      "Humidity of each city",
		}, []string{"location", "timestamp"}),
		Windspeed: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "gocast",
			Name:      "wind_speed",
			Help:      "Wind speed of each city",
		}, []string{"location", "timestamp"}),
		Pressure: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "gocast",
			Name:      "pressure",
			Help:      "Pressure of each city",
		}, []string{"location", "timestamp"}),
	}
	reg.MustRegister(m.Temperature)
	reg.MustRegister(m.Humidity)
	reg.MustRegister(m.Windspeed)
	reg.MustRegister(m.Pressure)
	return m
}

func RegisterMetrics(record model.WeatherRecord, hour int) {
	if len(record.Hours) != 0 {
		m.Temperature.With(prometheus.Labels{"location": record.Hours[hour].City, "timestamp": record.Hours[hour].TimeStamp}).Set(record.Hours[hour].Temperature)
		m.Humidity.With(prometheus.Labels{"location": record.Hours[hour].City, "timestamp": record.Hours[hour].TimeStamp}).Set(float64(record.Hours[hour].RelativeHumidity))
		m.Windspeed.With(prometheus.Labels{"location": record.Hours[hour].City, "timestamp": record.Hours[hour].TimeStamp}).Set(record.Hours[hour].WindSpeed)
		m.Pressure.With(prometheus.Labels{"location": record.Hours[hour].City, "timestamp": record.Hours[hour].TimeStamp}).Set(record.Hours[hour].PressureMSL)
	}
}

func RegisterDayMetrics(record model.WeatherRecord) {
	if len(record.Hours) != 0 {
		for i := 0; i < len(record.Hours); i++ {
			m.Temperature.With(prometheus.Labels{"location": record.Hours[i].City, "timestamp": record.Hours[i].TimeStamp}).Set(record.Hours[i].Temperature)
			m.Humidity.With(prometheus.Labels{"location": record.Hours[i].City, "timestamp": record.Hours[i].TimeStamp}).Set(float64(record.Hours[i].RelativeHumidity))
			m.Windspeed.With(prometheus.Labels{"location": record.Hours[i].City, "timestamp": record.Hours[i].TimeStamp}).Set(record.Hours[i].WindSpeed)
			m.Pressure.With(prometheus.Labels{"location": record.Hours[i].City, "timestamp": record.Hours[i].TimeStamp}).Set(record.Hours[i].PressureMSL)
		}
	}
}

func getMetrics() *model.FirstMetricStruct {
	if m != nil {
		return m
	}
	return nil
}

*/
