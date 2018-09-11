package db

type Playlist struct {
	PlaylistID string  `json:"playlistID"`
	Videos     []Video `json:"videos,omitempty"`
}

type Video struct {
	VideoTitle int    `json:"videoTitle"`
	VideoID    string `json:"videoID"`
}

type PlaylistPostRequest struct {
	ID string `json:"id"`
}
