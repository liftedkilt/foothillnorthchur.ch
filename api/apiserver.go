package main

import (
	"log"
	"net/http"

	"github.com/liftedkilt/foothillnorthchur.ch/api/api"
	"github.com/gorilla/mux"
)

const (
	port = ":8080"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/playlist", api.GetPlaylistLatest).Methods("GET")
	router.HandleFunc("/playlist/{week}", api.GetPlaylistByWeek).Methods("GET")
	router.HandleFunc("/playlist/{week}", api.CreatePlaylistByWeek).Methods("POST")
	router.HandleFunc("/songs/{title}", api.GetSongByTitle).Methods("GET")
	router.HandleFunc("/songs", api.GetSongs).Methods("GET")

	log.Println("Running API server on port", port)
	log.Fatal(http.ListenAndServe(port, router))
}

// GET /playlist/latest
// GET /playlist/{week}
// GET /songs/{title}
// GET /songs
