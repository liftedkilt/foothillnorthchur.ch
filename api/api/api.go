package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	// "github.com/liftedkilt/foothillnorthchur.ch/api/db"
	"github.com/liftedkilt/foothillnorthchur.ch/api/db"
)

// Playlist routes
// router.HandleFunc("/playlist", api.GetPlaylistLatest).Methods("GET")
// router.HandleFunc("/playlist/{date}", api.GetPlaylist).Methods("GET")
// router.HandleFunc("/playlist/{date}", api.CreatePlaylist).Methods("POST")

func GetPlaylistLatest(w http.ResponseWriter, r *http.Request) {
	// Calculate latest playlist
	// Retrieve playlist
	// json.NewEncoder(w).Encode(playlist)
	return
}

func GetPlaylist(w http.ResponseWriter, r *http.Request) {
	// params := mux.Vars(r)
	// date := params["date"]

	// playlist := db.GetPlaylistByDate(date)

	// json.NewEncoder(w).Encode(playlist)
	return
}

func CreatePlaylist(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	date := params["date"]

	log.Println("Creating playlist for date", date)

	// Parse playlist into struct
	var playlistRequest = new(db.PlaylistPostRequest)
	reqbytes, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqbytes, &playlistRequest)

	// Construct url and request playlist from parser
	url := "http://localhost:8080/playlist/" + playlistRequest.ID
	resp, err := http.Post(url, "", nil)

	if err != nil {
		fmt.Println(err)
		fmt.Println("ERROR: Failed to parse \"" + url + "\"")
		return
	}

	// Extract response from body
	b := resp.Body
	defer b.Close()
	resbytes, _ := ioutil.ReadAll(b)

	// Unmarshal response into []Video
	var parserResponse []db.Video
	json.Unmarshal(resbytes, &parserResponse)

	//  Initialize Playlist struct, and populate required fields
	var playlist = new(db.Playlist)
	playlist.PlaylistID = playlistRequest.ID
	playlist.Videos = parserResponse

	// db.CreatePlaylist(date, playlist)

	// json.NewEncoder(w).Encode(titles)
}

// Video routes
// router.HandleFunc("/video/{videoID}", api.GetVideo).Methods("GET")
// router.HandleFunc("/video/{videoID}", api.CreateVideo).Methods("POST")
// router.HandleFunc("/video/{videoID}", api.DeleteVideo).Methods("DELETE")
