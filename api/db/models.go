package db

type Playlist struct {
	PlaylistID string  `json:"playlistID"`
	Videos     []Video `json:"videos,omitempty"`
}

type Video struct {
	VideoID    string `json:"videoID"`
	VideoTitle string `json:"videoTitle"`
}
