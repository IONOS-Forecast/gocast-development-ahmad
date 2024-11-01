---
title: README.md
created: '2024-11-01T14:28:47.628Z'
modified: '2024-11-01T16:04:39.623Z'
---

README.md

### Gocast

---

A weather application built with goloang. The application only works for cities in Germany


|command line|environmental variable|default|description|
|------------|-----------|-------|-----------|
|-|`CITY_API_KEY`|-|API key the cities API, you can get yours at: https://openweathermap.org/api|
|-|`DB_USER`|forecast|database user for weather database|
|-|`DB_PASSWORD`|forecast|database password for weather database|
|-|`DB_NAME`|forecast|name of the weather database|
|-|`DB_ADDRESS`|127.0.0.1:5544|address for the weather database|
|-|`CITY`|-|city of the desired weather record|
|-|`HOUR`|-|hour of the desired weather record|
|-|`DAY`|-|day of the desired weather record|
|-|`MONTH`|-|month of the desired weather record|
|-|`YEAR`|-|year of the desired weather record|
|-|`N_MINUTES`|-|how often the weather data should be called in minutes (disabled)|
---
screenshots
---
---
### how to run:

#### stadard output:
run `go run main.go` on the main application directory after adjusting the enviromental variable values to your wish
#### with graphana:
run `sudo docker-compose up` on the main application directory
#### get weather data as json:

run `curl -s "http://127.0.0.1:8080/?date=2024-09-01&city=berlin" | jq -r '.weather[] | to_entries[] | "\(.key): \(.value)"'`

to get a specific metric value in a specific hour, you can use this command `curl -s "http://127.0.0.1:8080/?date=2024-09-01&city=berlin" | jq -r '.weather[<hour>].<key>'` change `hour` and `key`
example: `curl -s "http://127.0.0.1:8080/?date=2024-09-01&city=berlin" | jq -r '.weather[1].temperature'`


