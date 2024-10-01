package usecases_test

import (
	"testing"

	"github.com/Moreira-Henrique-Pedro/music-weather/src/usecases"
	"github.com/stretchr/testify/assert"
)

func TestDeterminePlaylistGenre_HotWeather(t *testing.T) {
	// Inicializa a classe a ser testada, nesse caso NewTemperatureUseCase
	tu := usecases.NewTemperatureUseCase()

	// Atribui um valor a função DeterminePlaylistGenre, nesse caso um valor acima de 25
	genre, err := tu.DeterminePlaylistGenre(30)
	assert.NoError(t, err)
	assert.Equal(t, usecases.MusicGenrePop, genre)
}

func TestDeterminePlaylistGenre_MildWeather(t *testing.T) {
	// Inicializa a classe a ser testada, nesse caso NewTemperatureUseCase
	tu := usecases.NewTemperatureUseCase()

	// Atribui um valor a função DeterminePlaylistGenre, nesse caso um valor entre 10 e 25
	genre, err := tu.DeterminePlaylistGenre(15)
	assert.NoError(t, err)
	assert.Equal(t, usecases.MusicGenreRock, genre)
}

func TestDeterminePlaylistGenre_ColdWeather(t *testing.T) {
	// Inicializa a classe a ser testada, nesse caso NewTemperatureUseCase
	tu := usecases.NewTemperatureUseCase()

	// Atribui um valor a função DeterminePlaylistGenre, nesse caso um valor abaixo de 10
	genre, err := tu.DeterminePlaylistGenre(5)
	assert.NoError(t, err)
	assert.Equal(t, usecases.MusicGenreClassical, genre)
}
