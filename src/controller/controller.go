package controller

import (
	"fmt"
	"net/http"

	"github.com/Moreira-Henrique-Pedro/music-weather/src/model"
	"github.com/Moreira-Henrique-Pedro/music-weather/src/service"
	"github.com/Moreira-Henrique-Pedro/music-weather/src/usecases"
	"github.com/gin-gonic/gin"
)

type LocationController struct {
	weatherService *service.WeatherService
	spotifyService *service.SpotifyService
	usecases       *usecases.TemperatureUseCase
}

func NewLocationController(weatherService *service.WeatherService, spotifyService *service.SpotifyService, usecases *usecases.TemperatureUseCase) *LocationController {
	return &LocationController{
		weatherService: weatherService,
		spotifyService: spotifyService,
		usecases:       usecases,
	}
}

func (c *LocationController) InitRoutes() {

	app := gin.Default()

	app.POST("/music_weather", c.handleFunc)

	app.Run(":8000")

}

func (c *LocationController) handleFunc(ctx *gin.Context) {
	location := new(model.Location)
	if err := ctx.ShouldBindJSON(&location); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": err},
		)
		return
	}

	temp, err := c.weatherService.GetWeather(*location)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err},
		)
		return
	}

	fmt.Printf("Sua temperatura é %v", temp)

	genre, err := c.usecases.DeterminePlaylistGenre(temp)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"error": "Failed to determine playlist genre", "details": err.Error()},
		)
		return
	}

	playlist, err := c.spotifyService.GetPlaylistByGenre(genre)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"error": "error to get playlist", "details": err.Error()},
		)
		return
	}

	ctx.JSON(
		http.StatusOK,
		gin.H{"ID": playlist.ID,
			"Name":        playlist.Name,
			"URL":         playlist.URL,
			"ImageURL":    playlist.ImageURL,
			"TrackCount":  playlist.TrackCount,
			"City":        location,
			"Temperature": temp,
			"Genre":       genre,
		},
	)

}
