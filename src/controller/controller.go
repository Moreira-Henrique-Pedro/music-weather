package controller

import (
	"fmt"
	"net/http"

	"github.com/Moreira-Henrique-Pedro/music-weather/src/model"
	"github.com/Moreira-Henrique-Pedro/music-weather/src/service"
	"github.com/gin-gonic/gin"
)

type LocationController struct {
	service *service.WeatherService
}

func NewLocationController(service *service.WeatherService) *LocationController {
	return &LocationController{
		service: service,
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

	temp, err := c.service.GetWeather(*location)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err},
		)
		return
	}

	fmt.Printf("Sua temperatura é %v", temp)

}
