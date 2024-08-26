package service

import (
	"context"
	"math/rand"
	"os"
	"time"

	"github.com/Moreira-Henrique-Pedro/music-weather/src/model"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
)

type SpotifyService struct {
	Client *spotify.Client
}

func NewSpotifyService() *SpotifyService {
	config := &clientcredentials.Config{
		ClientID:     os.Getenv("SPOTIFY_CLIENT_ID"),
		ClientSecret: os.Getenv("SPOTIFY_CLIENT_SECRET"),
		TokenURL:     spotify.TokenURL,
	}

	httpClient := config.Client(context.Background())
	client := spotify.NewClient(httpClient)

	return &SpotifyService{
		Client: &client,
	}
}

func (s *SpotifyService) GetPlaylistByGenre(genre string) (model.Playlist, error) {
	searchResult, err := s.Client.Search("genre:"+genre, spotify.SearchTypePlaylist)
	if err != nil {
		return model.Playlist{}, err
	}

	if len(searchResult.Playlists.Playlists) > 0 {

		rand.Seed(time.Now().UnixNano())

		randomIndex := rand.Intn(len(searchResult.Playlists.Playlists))

		playlist := searchResult.Playlists.Playlists[randomIndex]
		return convertSpotifyPlaylist(playlist), nil
	}

	return model.Playlist{}, nil
}

func convertSpotifyPlaylist(playlist spotify.SimplePlaylist) model.Playlist {
	return model.Playlist{
		ID:         playlist.SnapshotID,
		Name:       playlist.Name,
		URL:        playlist.ExternalURLs["spotify"],
		ImageURL:   playlist.Images[0].URL,
		TrackCount: int(playlist.Tracks.Total),
	}
}
