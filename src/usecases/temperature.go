package usecases

import "fmt"

type TemperatureUseCase struct {
}

func NewTemperatureUseCase() *TemperatureUseCase {
	return &TemperatureUseCase{}
}

func (tu *TemperatureUseCase) DeterminePlaylistGenre(temp float64) (string, error) {

	switch {

	case temp > 25:
		return MusicGenrePop, nil

	case temp >= 10 && temp <= 25:
		return MusicGenreRock, nil

	case temp < 10:
		return MusicGenreClassical, nil

	default:
		return "", fmt.Errorf("invalid temperature value: %v", temp)
	}
}
