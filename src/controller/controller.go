package controller

import (
	"fmt"
	"net/http"

	_ "github.com/Moreira-Henrique-Pedro/music-weather/docs"
	"github.com/Moreira-Henrique-Pedro/music-weather/src/model"
	"github.com/Moreira-Henrique-Pedro/music-weather/src/service"
	"github.com/Moreira-Henrique-Pedro/music-weather/src/usecases"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	// Rota para o endpoint da API
	app.POST("/music_weather", c.handleFunc)

	// Rota para a documentação Swagger
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	app.Run(":8000")

}

// @BasePath /api/v1

// @Summary Get music-weather
// @Description API criada para sugerir uma playlist no spotify com base na temperatura atual da cidade escolhida.
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

	fmt.Printf("Temp founded %v", temp)

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
