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

// @Summary Get music-weather
// @Description Get weather based on location and suggest a playlist.
// @Tags Music Weather
// @Accept  json
// @Produce  json
// @Param   location body model.Location true "Location Info"
// @Success 200 {object} model.Playlist
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /music_weather [post]
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
			"City":        location.City,
			"Temperature": temp,
			"Genre":       genre,
		},
	)

}
