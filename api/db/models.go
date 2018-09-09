package db

type Song struct {
	Title        string `json:"title"`
	Lyrics       string `json:"lyrics"`
	YoutubeTitle string `json:"youtubeTitle"`
}

// "_id": 20180909,
// "playlistID": asdfljansdzffnaf
// "playlist": [
// 	{
// 		"order": 1,
// 		"videoID": "aawofiacanwea",
// 	},
// 	{
// 		"order": 2,
// 		"videoID": "srtvsdtvzsetv",
// 	}]
// }

type Playlist struct {
	PlaylistID string  `json:"playlistID"`
	Videos     []Video `json:"videos,omitempty"`
}

type Video struct {
	Order   int    `json:"order"`
	VideoID string `json:"videoID"`
}

type PlaylistPostRequest struct {
	ID string `json:"id"`
}
