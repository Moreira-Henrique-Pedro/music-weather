package main

import (
	"github.com/Moreira-Henrique-Pedro/music-weather/src/controller"
	"github.com/Moreira-Henrique-Pedro/music-weather/src/service"
)

func main() {

	weatherService := service.NewWeatherService()
	locationController := controller.NewLocationController(weatherService)

	locationController.InitRoutes()
}
