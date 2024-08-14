package main

import (
	"log"

	"github.com/Moreira-Henrique-Pedro/music-weather/src/controller"
	"github.com/Moreira-Henrique-Pedro/music-weather/src/service"
	"github.com/Moreira-Henrique-Pedro/music-weather/src/usecases"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {

	tempUsecase := usecases.NewTemperatureUseCase()
	weatherService := service.NewWeatherService()
	spotifyService := service.NewSpotifyService()

	locationController := controller.NewLocationController(weatherService, spotifyService, tempUsecase)

	locationController.InitRoutes()
}
