package api

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"math/rand"

	"github.com/IONOS-Forecast/gocast-development-ahmad/pkg/model"
	"github.com/joho/godotenv"
)

const CityAPIKey = ""

var GetWeatherDataFromAPIStructs = []struct {
	name                     string
	date                     string
	citynumbers              []model.CityData
	expected_weather_records model.WeatherDataForDay
}{}

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

var CityInfoStructs = []struct {
	name string
	lat  float64
	long float64
}{
	{"berlin", 52.5170365, 13.3888599},
	{"hamburg", 53.550341, 10.000654},
	{"munich", 48.1371079, 11.5753822},
	{"cologne", 50.938361, 6.959974},
	{"frankfurt", 50.1106444, 8.6820917},
	{"stuttgart", 48.7784485, 9.1800132},
	{"dusseldorf", 51.2254018, 6.7763137},
	{"dortmund", 51.5142273, 7.4652789},
	{"essen", 51.4582235, 7.0158171},
	{"bremen", 53.0758196, 8.8071646},
	{"hanover", 52.3744779, 9.7385532},
	{"nuremberg", 49.453872, 11.077298},
	{"dresden", 51.0493286, 13.7381437},
	{"leipzig", 51.3406321, 12.3747329},
	{"chemnitz", 50.8322608, 12.9252977},
	{"mannheim", 49.4892913, 8.4673098},
	{"augsburg", 48.3668041, 10.8986971},
	{"wiesbaden", 50.0820384, 8.2416556},
	{"karlsruhe", 49.0068705, 8.4034195},
	{"freiburg im breisgau", 47.9960901, 7.8494005},
	{"regensburg", 49.0195333, 12.0974869},
	{"erfurt", 50.9777974, 11.0287364},
	{"mainz", 50.0012314, 8.2762513},
	{"rostock", 54.0924445, 12.1286127},
	{"kiel", 54.3227085, 10.135555},
	{"köthen (anhalt)", 51.751033, 11.973698},
	{"remscheid", 51.1798706, 7.1943544},
	{"moers", 51.451283, 6.62843},
	{"ulm", 48.3974003, 9.9934336},
	{"bielefeld", 52.0191005, 8.531007},
	{"jena", 50.9281717, 11.5879359},
	{"zwickau", 50.7185043, 12.4939267},
	{"gelsenkirchen", 51.5110321, 7.0960124},
	{"schwerin", 53.6288297, 11.4148038},
	{"siegen", 50.8749804, 8.0227233},
	{"flensburg", 54.7833021, 9.4333264},
	{"göttingen", 51.5328328, 9.9351811},
	{"reutlingen", 48.4919508, 9.2114144},
	{"trier", 49.7596208, 6.6441878},
	{"offenbach am main", 50.1055002, 8.7610698},
	{"ravensburg", 47.7811014, 9.612468},
	{"neuss", 51.1981778, 6.6916476},
	{"viersen", 51.2562118, 6.3905476},
	{"lübeck", 53.866444, 10.684738},
	{"leverkusen", 51.0324743, 6.9881194},
	{"iserlohn", 51.3746778, 7.6999713},
	{"marburg", 50.8090106, 8.7704695},
	{"norderstedt", 53.7089898, 9.9891914},
	{"haßloch", 49.362976, 8.2565755},
	{"bottrop", 51.521581, 6.929204},
	{"göppingen", 48.7031377, 9.6541116},
	{"herne", 51.5380394, 7.219985},
	{"jülich", 50.9220931, 6.3611015},
	{"tübingen", 48.5236164, 9.0535531},
	{"cottbus - chóśebuz", 51.7567447, 14.3357307},
	{"untere mühle", 48.035558, 8.8905466},
	{"landshut", 48.536217, 12.1516551},
	{"krefeld", 51.3331205, 6.5623343},
	{"passau", 48.5748229, 13.4609744},
	{"rheinberg", 51.5458979, 6.6014097},
	{"bad salzuflen", 52.0771518, 8.7554739},
	{"lörrach", 47.6120896, 7.6607218},
	{"wismar", 53.8909832, 11.4647932},
	{"albstadt", 48.233048, 8.9991483},
	{"freudenstadt", 48.4637727, 8.4111727},
	{"grevenbroich", 51.0905783, 6.5835365},
	{"brunswick", 52.2646577, 10.5236066},
	{"aalen", 48.8362705, 10.0931765},
	{"herford", 52.1152245, 8.6711118},
	{"würzburg", 49.79245, 9.932966},
	{"darmstadt", 49.8851869, 8.6736295},
	{"schwäbisch gmünd", 48.7999036, 9.7977584},
	{"völklingen", 49.2522866, 6.859519},
	{"delmenhorst", 53.048095, 8.6286066},
	{"paderborn", 51.71895955, 8.764869778177559},
	{"saarbrücken", 49.234362, 6.996379},
	{"lünen", 51.6142482, 7.5228088},
	{"schweinfurt", 50.0499945, 10.233302},
	{"bayreuth", 49.9427202, 11.5763079},
	{"hanau", 50.132881, 8.9169797},
	{"friedrichshafen", 47.6500279, 9.4800858},
	{"rheinbach", 50.6256808, 6.9491436},
	{"filderstadt", 48.6616311, 9.2236805},
	{"fürth", 49.4772475, 10.9893626},
	{"freising", 48.4008273, 11.7439565},
	{"bad honnef", 50.6448582, 7.2272534},
	{"geilenkirchen", 50.963605, 6.1199802},
	{"schwetzingen", 49.3832919, 8.5735135},
	{"hagen", 51.3582945, 7.473296},
	{"alsdorf", 49.8885782, 6.4643156},
	{"neumünster", 54.0757442, 9.9815377},
	{"strausberg", 52.5814009, 13.8833952},
	{"neubrandenburg", 53.5574458, 13.2602781},
	{"waren (müritz)", 53.5156249, 12.6850606},
	{"kuckssee", 53.5341859, 13.089188446848897},
	{"oldenburg", 53.1389753, 8.2146017},
	{"mönchengladbach", 51.1946983, 6.4353641},
	{"heidelberg", 49.4093582, 8.694724},
	{"ingolstadt", 48.7630165, 11.4250395},
	{"kassel", 51.3154546, 9.4924096},
}

func TestGetCityData(t *testing.T) {

	var cityinfo model.CityData

	for i := 0; i < len(CityInfoStructs); i++ {
		t.Run(CityInfoStructs[i].name, func(t *testing.T) {
			cityinfo = GetCityData(CityInfoStructs[i].name, CityAPIKey)
			if CityInfoStructs[i].lat != cityinfo.Lat || CityInfoStructs[i].long != cityinfo.Lon {
				t.Error(fmt.Printf("Test failed: input {\"%s\",%v,%v} output{\"%s\",%v,%v}", CityInfoStructs[i].name, CityInfoStructs[i].lat, CityInfoStructs[i].long, cityinfo.Name, cityinfo.Lat, cityinfo.Lon))
			}
		})
	}
}

func TestCreateCityInfoStructs(t *testing.T) {

	var cityinfo model.CityData

	godotenv.Load()
	CityAPIKey := os.Getenv("CITY_API_KEY")

	for i := 0; i < len(cities); i++ {
		cityinfo = GetCityData(cities[i], CityAPIKey)
		fmt.Println(fmt.Sprintf("{%s,%f,%f}", cityinfo.Name, cityinfo.Lat, cityinfo.Lon))
	}

}

func createRandomDates(howmanydates int) []string {

	if howmanydates < 1 {
		log.Fatal("number or dates cant be under 1")
	}

	var date string
	var dates []string

	for i := 0; i < howmanydates-1; i++ {

		rand.Seed(time.Now().UnixNano())

		// Define the start date as January 1, 2017
		startDate := time.Date(2010, 1, 1, 0, 0, 0, 0, time.UTC)

		// Get yesterday's date
		yesterday := time.Now().AddDate(0, 0, -1).Truncate(24 * time.Hour)

		// Calculate the duration between startDate and yesterday
		duration := yesterday.Sub(startDate)

		// Generate a random duration within the range
		randomDuration := time.Duration(rand.Int63n(int64(duration)))

		// Calculate the random date by adding the random duration to the start date
		randomDate := startDate.Add(randomDuration)

		// Print the random date
		date = randomDate.Format("2006-01-02")
		dates[i] = date
		fmt.Println(fmt.Sprintf(date))
	}

	return dates
}
