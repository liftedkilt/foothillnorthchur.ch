package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/liftedkilt/foothillnorthchur.ch/parser/parser"
)

func GetPlaylistLatest(w http.ResponseWriter, r *http.Request) {
	url := "https://youtube.com/playlist?list=" + params["playlistID"]

	log.Println("Processing playlist:", url)

	titles := parser.Parse(url)
	log.Printf("Playlist processed: %+v\n", titles)

	json.NewEncoder(w).Encode(titles)
}

func GetPlaylistByWeek(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	week := params["week"]

	log.Println("Processing playlist:", url)

	titles := parser.Parse(url)
	log.Printf("Playlist processed: %+v\n", titles)

	json.NewEncoder(w).Encode(titles)
}

func GetSongByTitle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	url := "https://youtube.com/playlist?list=" + params["playlistID"]

	log.Println("Processing playlist:", url)

	titles := parser.Parse(url)
	log.Printf("Playlist processed: %+v\n", titles)

	json.NewEncoder(w).Encode(titles)
}

func GetSongs(w http.ResponseWriter, r *http.Request) {
	log.Println("Processing playlist:", url)

	titles := parser.Parse(url)
	log.Printf("Playlist processed: %+v\n", titles)

	json.NewEncoder(w).Encode(titles)
}

// router.HandleFunc("/playlist", api.GetPlaylistLatest).Methods("GET")
// router.HandleFunc("/playlist/{week}", api.GetPlaylistByWeek).Methods("GET")
// router.HandleFunc("/songs/{title}", api.GetSongByTitle).Methods("GET")
// router.HandleFunc("/songs", api.GetSongs).Methods("GET")
