package server

import (
	"github.com/IONOS-Forecast/gocast-development-ahmad/pkg/model"
	"github.com/prometheus/client_golang/prometheus"
)

var m *model.FirstMetricStruct

func FirstMetric(reg prometheus.Registerer) *model.FirstMetricStruct {
	m := &model.FirstMetricStruct{
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

func RegisterFirstMetric(bla model.WeatherDataForDay) {

	layout := "2006-01-02 15:04:05-07"

	for i := 0; i < len(bla.WeatherDataForTheDay); i++ {
		m.Temperature.With(prometheus.Labels{"location": bla.WeatherDataForTheDay[i].City, "timestamp": bla.WeatherDataForTheDay[i].TimeStamp.Format(layout)}).Set(bla.WeatherDataForTheDay[i].Temperature)
		m.Humidity.With(prometheus.Labels{"location": bla.WeatherDataForTheDay[i].City, "timestamp": bla.WeatherDataForTheDay[i].TimeStamp.Format(layout)}).Set(float64(bla.WeatherDataForTheDay[i].RelativeHumidity))
		m.Windspeed.With(prometheus.Labels{"location": bla.WeatherDataForTheDay[i].City, "timestamp": bla.WeatherDataForTheDay[i].TimeStamp.Format(layout)}).Set(bla.WeatherDataForTheDay[i].WindSpeed)
		m.Pressure.With(prometheus.Labels{"location": bla.WeatherDataForTheDay[i].City, "timestamp": bla.WeatherDataForTheDay[i].TimeStamp.Format(layout)}).Set(bla.WeatherDataForTheDay[i].PressureMsl)
	}

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
