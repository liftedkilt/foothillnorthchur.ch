package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/liftedkilt/foothillnorthchur.ch/api/api"
)

const (
	port = ":8081"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/playlist", api.GetPlaylistLatest).Methods("GET")

	router.HandleFunc("/playlist/{date}", api.GetPlaylist).Methods("GET")
	router.HandleFunc("/playlist/{date}", api.CreatePlaylist).Methods("POST")

	log.Println("Running API server on port", port)

	log.Fatal(http.ListenAndServe(port, router))
}
