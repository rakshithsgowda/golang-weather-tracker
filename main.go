package main

import (
	"fmt"
	"weather_tracker/router"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("file can't be found", err)
		return
	}
	router.Routing()
}
