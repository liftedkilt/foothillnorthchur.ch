package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/liftedkilt/foothillnorthchur.ch/parser/parser"
)

const (
	port = ":8080"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/playlist/{playlistID}", parsePlaylist).Methods("POST")

	log.Println("Running Parser server on port", port)
	log.Fatal(http.ListenAndServe(port, router))
}

func parsePlaylist(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	url := "https://youtube.com/playlist?list=" + params["playlistID"]

	log.Println("Processing playlist:", url)

	titles := parser.Parse(url)
	log.Printf("Playlist processed: %+v\n", titles)

	json.NewEncoder(w).Encode(titles)

}
