package router

import (
	"net/http"
	"weather_tracker/controller"
)

var WeatherRoutes = Routes{
	Route{"GEt temp of perticular city", http.MethodGet, "/getWeather", controller.GetCityWeather},
}
