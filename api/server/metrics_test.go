package server

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"math/rand"

	"github.com/IONOS-Forecast/gocast-development-ahmad/pkg/model"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/expfmt"
)

var cities = []string{
	"berlin",
	"hamburg",
	"munich",
	"cologne",
	"frankfurt",
	"stuttgart",
	"dusseldorf",
	"dortmund",
	"essen",
	"bremen",
	"hanover",
	"nuremberg",
	"dresden",
	"leipzig",
	"chemnitz",
	"mannheim",
	"augsburg",
	"wiesbaden",
	"karlsruhe",
	"freiburg im breisgau",
	"regensburg",
	"erfurt",
	"mainz",
	"rostock",
	"kiel",
	"köthen (anhalt)",
	"remscheid",
	"moers",
	"ulm",
	"bielefeld",
	"jena",
	"zwickau",
	"gelsenkirchen",
	"schwerin",
	"siegen",
	"flensburg",
	"göttingen",
	"reutlingen",
	"trier",
	"offenbach am main",
	"ravensburg",
	"neuss",
	"viersen",
	"lübeck",
	"leverkusen",
	"iserlohn",
	"marburg",
	"norderstedt",
	"haßloch",
	"bottrop",
	"göppingen",
	"herne",
	"jülich",
	"tübingen",
	"cottbus - chóśebuz",
	"untere mühle",
	"landshut",
	"krefeld",
	"passau",
	"rheinberg",
	"bad salzuflen",
	"lörrach",
	"wismar",
	"albstadt",
	"freudenstadt",
	"grevenbroich",
	"brunswick",
	"aalen",
	"herford",
	"würzburg",
	"darmstadt",
	"schwäbisch gmünd",
	"völklingen",
	"delmenhorst",
	"paderborn",
	"saarbrücken",
	"lünen",
	"schweinfurt",
	"bayreuth",
	"hanau",
	"friedrichshafen",
	"rheinbach",
	"filderstadt",
	"fürth",
	"freising",
	"bad honnef",
	"geilenkirchen",
	"schwetzingen",
	"hagen",
	"alsdorf",
	"neumünster",
	"strausberg",
	"neubrandenburg",
	"waren",
	"kuckssee",
	"oldenburg",
	"mönchengladbach",
	"heidelberg",
	"ingolstadt",
	"kassel",
}

var MetricsValuesStruct = []struct {
	name        string
	temperature float64
	humidity    int
	windspeed   float64
	pressure    float64
}{
	{"test1", 22.5, 55, 5.2, 1013.1},
	{"test2", 25.3, 60, 7.8, 1012.5},
	{"test3", 19.8, 70, 3.5, 1010.8},
	{"test4", 30.1, 45, 10.0, 1011.6},
	{"test5", 28.4, 50, 8.2, 1012.3},
	{"test6", 21.7, 65, 6.4, 1014.2},
	{"test7", 23.9, 58, 4.9, 1012.0},
	{"test8", 26.8, 62, 9.1, 1010.0},
	{"test9", 24.0, 57, 5.6, 1013.0},
	{"test10", 20.5, 75, 2.3, 1009.5},
	{"test11", 27.1, 49, 8.5, 1011.2},
	{"test12", 29.0, 55, 11.3, 1013.5},
	{"test13", 22.8, 61, 6.0, 1014.8},
	{"test14", 19.2, 68, 3.2, 1008.7},
	{"test15", 31.0, 42, 12.0, 1012.6},
	{"test16", 26.5, 53, 7.9, 1011.9},
	{"test17", 23.1, 60, 5.5, 1013.4},
	{"test18", 25.9, 64, 6.7, 1012.1},
	{"test19", 20.3, 72, 2.8, 1009.9},
	{"test20", 28.2, 48, 9.4, 1011.5},
	{"test21", 21.5, 66, 5.1, 1013.3},
	{"test22", 29.5, 54, 10.4, 1012.9},
	{"test23", 24.6, 58, 7.3, 1010.6},
	{"test24", 22.0, 59, 6.2, 1014.1},
	{"test25", 19.9, 71, 3.8, 1008.4},
	{"test26", 27.8, 44, 11.1, 1012.7},
	{"test27", 30.2, 50, 12.5, 1013.2},
	{"test28", 23.4, 55, 6.9, 1011.7},
	{"test29", 26.7, 63, 8.1, 1012.8},
	{"test30", 20.1, 69, 2.5, 1009.0},
}

var timestamps = []string{
	"2021-02-28T00:00:00Z",
	"2019-07-04T00:00:00Z",
	"2025-12-19T00:00:00Z",
	"2030-01-01T00:00:00Z",
	"1980-11-15T00:00:00Z",
	"1965-05-25T00:00:00Z",
	"2028-08-09T00:00:00Z",
	"2022-10-31T00:00:00Z",
	"2026-03-14T00:00:00Z",
	"2015-09-05T00:00:00Z",
	"1930-02-28T00:00:00Z",
	"1960-12-31T00:00:00Z",
	"1999-07-01T00:00:00Z",
	"2012-02-29T00:00:00Z",
	"2050-02-28T00:00:00Z",
	"2100-02-28T00:00:00Z",
	"2200-02-28T00:00:00Z",
	"2400-02-29T00:00:00Z",
	"8000-08-15T00:00:00Z",
	"3000-01-01T00:00:00Z",
	"1800-07-04T00:00:00Z",
	"1776-07-04T00:00:00Z",
	"2540-11-11T00:00:00Z",
	"1900-02-28T00:00:00Z",
	"2023-03-01T00:00:00Z",
	"2055-05-30T00:00:00Z",
	"2301-12-25T00:00:00Z",
	"9999-11-11T00:00:00Z",
	"2101-03-15T00:00:00Z",
	"2052-02-29T00:00:00Z",
}

func TestBla(t *testing.T) {

	fmt.Println("ew") // add whatever so a new test can be started

	weatherinfo := model.WeatherDataForDay{
		WeatherDataForTheDay: make([]model.Weather_record, 24),
	}

	reg := prometheus.NewRegistry()
	FirstMetric(reg)
	PrometeusMux(reg)

	for i := 0; i < len(MetricsValuesStruct); i++ {
		t.Run(MetricsValuesStruct[i].name, func(t *testing.T) {

			hour := rand.Intn(24)
			timestampRandom := rand.Intn(len(timestamps))
			cityRandom := rand.Intn(len(cities))
			tempCity := cities[cityRandom]

			timestampInTime, _ := time.Parse(time.RFC3339, timestamps[timestampRandom])
			timestampInTime = timestampInTime.Add(time.Duration(hour) * time.Hour)
			tempTimestamp := timestampInTime.Format(time.RFC3339)

			weatherinfo.WeatherDataForTheDay[hour].TimeStamp = timestampInTime
			weatherinfo.WeatherDataForTheDay[hour].Temperature = MetricsValuesStruct[i].temperature
			weatherinfo.WeatherDataForTheDay[hour].RelativeHumidity = MetricsValuesStruct[i].humidity
			weatherinfo.WeatherDataForTheDay[hour].WindSpeed = MetricsValuesStruct[i].windspeed
			weatherinfo.WeatherDataForTheDay[hour].PressureMsl = MetricsValuesStruct[i].pressure
			weatherinfo.WeatherDataForTheDay[hour].City = tempCity

			SetValueFirstMetric(weatherinfo, hour)

			outputTemperature, err := getMetrics("gocast_temperature", tempTimestamp, tempCity)
			if err != nil {
				t.Error("temperature couldnt be imported via http")
			}
			outputHumidity, err := getMetrics("gocast_humidity", tempTimestamp, tempCity)
			if err != nil {
				t.Error("humidity couldnt be imported via http")
			}
			outputWindspeed, err := getMetrics("gocast_wind_speed", tempTimestamp, tempCity)
			if err != nil {
				t.Error("windspeed couldnt be imported via http")
			}
			outputPressure, err := getMetrics("gocast_pressure", tempTimestamp, tempCity)
			if err != nil {
				t.Error("pressure couldnt be imported via http")
			}

			if MetricsValuesStruct[i].temperature != outputTemperature || MetricsValuesStruct[i].humidity != int(outputHumidity) || MetricsValuesStruct[i].windspeed != outputWindspeed || MetricsValuesStruct[i].pressure != outputPressure {
				t.Errorf("Test failed\n input  temperature:%v humidity:%v windspeed:%v pressure:%v\n output temperature:%v humidity:%v windspeed:%v pressure:%v\n", MetricsValuesStruct[i].temperature, MetricsValuesStruct[i].humidity, MetricsValuesStruct[i].windspeed, MetricsValuesStruct[i].pressure, outputTemperature, outputHumidity, outputWindspeed, outputPressure)
			}

		})
	}
	/* select {} */

}

func getMetrics(metricName, timestamp, city string) (float64, error) {
	resp, err := http.Get("http://localhost:8081/metrics")
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var parser expfmt.TextParser
	metrics, err := parser.TextToMetricFamilies(resp.Body) // a map of prometeus MetricFamily struct
	if err != nil {
		return 0, err
	}

	if metricFamily, ok := metrics[metricName]; ok { // 'metricFamily' is a map of of prometeus MetricFamily struct of a specific metric (ex. temperature)
		for _, metric := range metricFamily.GetMetric() { // a struct of metrics with labels (name, value), and a gauge value
			labelMap := make(map[string]string)
			for _, label := range metric.GetLabel() { // labels that exist in the metric, each metric has one or more labels
				labelMap[label.GetName()] = label.GetValue()

				if labelMap["timestamp"] == timestamp && labelMap["location"] == city {
					return metric.GetGauge().GetValue(), nil
				}
			}
		}
	}

	return 0, fmt.Errorf("metrics found no values!")
}

func PrometeusMux(reg *prometheus.Registry) { // learn

	promHandler := promhttp.HandlerFor(reg, promhttp.HandlerOpts{})
	pMux := http.NewServeMux()
	pMux.Handle("/metrics", promHandler)

	go func() {
		http.ListenAndServe(":8081", pMux)
	}()

}
