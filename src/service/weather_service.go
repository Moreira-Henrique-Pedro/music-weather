package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/Moreira-Henrique-Pedro/music-weather/src/model"
)

type WeatherService struct {
	//
}

func NewWeatherService() *WeatherService {
	return &WeatherService{}
}

func (s *WeatherService) GetWeather(location model.Location) (float64, error) {

	weather := new(model.WeatherResponse)

	apiKey := os.Getenv("OPENWEATHER_API_KEY")
	if apiKey == "" {
		return 0, fmt.Errorf("API key not found in environment variables")
	}

	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", url.QueryEscape(location.City), apiKey)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making the request:", err)
		return 0, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Request failed with status:", resp.StatusCode)
		return 0, fmt.Errorf("request failed with status %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading the response body:", err)
		return 0, err
	}

	err = json.Unmarshal(body, weather)
	if err != nil {
		fmt.Println("Error unmarshaling the JSON:", err)
		return 0, err
	}

	return weather.Main.Temp, nil

}
