package main

import (
	"log"

	"github.com/Moreira-Henrique-Pedro/music-weather/src/controller"
	"github.com/Moreira-Henrique-Pedro/music-weather/src/service"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {

	weatherService := service.NewWeatherService()
	locationController := controller.NewLocationController(weatherService)

	locationController.InitRoutes()
}
