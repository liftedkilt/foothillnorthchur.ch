package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/liftedkilt/foothillnorthchur.ch/api/api"
)

const (
	port = ":8888"
)

func main() {
	router := mux.NewRouter()

	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With"})
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})

	router.HandleFunc("/playlist", api.GetPlaylistLatest).Methods("GET")

	router.HandleFunc("/playlist/{date}", api.GetPlaylist).Methods("GET")
	router.HandleFunc("/playlist/{date}/{playlistID}", api.CreatePlaylist).Methods("POST")

	log.Println("Running API server on port", port)

	log.Fatal(http.ListenAndServe(port, handlers.CORS(allowedHeaders, allowedOrigins, allowedMethods)(router)))
}
