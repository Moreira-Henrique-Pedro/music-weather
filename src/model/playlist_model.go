package model

type Playlist struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	URL         string  `json:"url"`
	ImageURL    string  `json:"image_url"`
	TrackCount  int     `json:"track_count"`
	City        string  `json:"city"`
	Temperature float64 `json:"temperature"`
	Genre       string  `json:"genre"`
}
