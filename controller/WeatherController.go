package controller

import (
	"encoding/json"
	"net/http"
	"weather_tracker/helper"

	"github.com/gin-gonic/gin"
)

func GetCityWeather(c *gin.Context) {
	city := c.Query("city")
	if city == "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "city can't be empty",
		})
		return
	}
	resp, err := helper.CallWeatherApi(city)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "there is something wrong happening",
		})
		return
	}
	defer resp.Body.Close()

	var d map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&d)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data": map[string]interface{}{
			"city":         city,
			"city_weather": d,
		},
	})
}
