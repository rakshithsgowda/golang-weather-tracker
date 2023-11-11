package helper

import (
	"net/http"
	"os"
)

func CallWeatherApi(city string) (*http.Response, error) {
	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?APPID=" + os.Getenv("APIKEY") + "&q=" + city)
	return resp, err
}
