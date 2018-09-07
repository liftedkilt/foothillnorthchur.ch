package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-xorm/xorm"
	"github.com/julienschmidt/httprouter"
	"github.com/liftedkilt/foothillnorthchur.ch/db/model"
)

// Server is an http server that handles REST requests.
type Server struct {
	db *xorm.Engine
}

// NewServer creates a new instance of a Server.
func NewServer(db *xorm.Engine) *Server {
	return &Server{db: db}
}

// RegisterRouter registers a router onto the Server.
func (s *Server) RegisterRouter(router *httprouter.Router) {
	router.GET("/ping", s.ping)

	router.GET("/songs", s.getSongs)
	router.POST("/songs", s.createSong)
	router.GET("/songs/:songID", s.getSongsByID)
	router.DELETE("/songs/:songID", s.deleteSong)

	router.GET("/playlist", s.getLatestPlaylist)
	router.POST("/playlist", s.createPlaylist)
	// router.DELETE("/playlist/:playlistID", s.deleteSong)

}

func (s *Server) ping(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	writeTextResult(w, "go/xorm")
}

func (s *Server) getSongs(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var songs []model.Songs
	s.db.Find(&songs)
	writeJSONResult(w, songs)
	return
}

func (s *Server) getSongsByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	songID := ps.ByName("songID")
	songIDi, _ := strconv.Atoi(songID)
	songIDi64 := int64(songIDi)

	var song = model.Songs{Id: songIDi64}
	_, err := s.db.Get(&song)
	if err != nil {
		http.Error(w, err.Error(), errToStatusCode(err))
	} else {
		writeJSONResult(w, song)
	}
	return

}

func (s *Server) createSong(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var song model.Songs
	if err := json.NewDecoder(r.Body).Decode(&song); err != nil {
		http.Error(w, err.Error(), errToStatusCode(err))
		return
	}

	affected, err := s.db.Insert(song)
	if err != nil {
		http.Error(w, err.Error(), errToStatusCode(err))
	} else {
		writeJSONResult(w, affected)
	}

}

func (s *Server) deleteSong(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	songID := ps.ByName("songID")
	songIDint, _ := strconv.Atoi(songID)
	affected, err := s.db.Id(songIDint).Delete(&model.Songs{})
	if err != nil {
		http.Error(w, err.Error(), errToStatusCode(err))
	} else {
		writeJSONResult(w, affected)
	}
}

//
// Playlist Endpoints
//

func (s *Server) getLatestPlaylist(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var playlist []model.Playlist

	s.db.Desc("Id").Limit(1).Find(&playlist)
	writeJSON(w, playlist)
	writeJSONResult(w, playlist)
	return
}

func (s *Server) createPlaylist(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var playlist model.Playlist

	if err := json.NewDecoder(r.Body).Decode(&playlist); err != nil {
		http.Error(w, err.Error(), errToStatusCode(err))
		return
	}

	affected, err := s.db.Insert(playlist)
	if err != nil {
		http.Error(w, err.Error(), errToStatusCode(err))
	} else {
		writeJSONResult(w, affected)
	}

}

//
//  Helper Functions
//

func writeTextResult(w http.ResponseWriter, res string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, res)
}

func writeJSONResult(w http.ResponseWriter, res interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		panic(err)
	}
}

func writeJSON(w http.ResponseWriter, res interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.Marshal(res); err != nil {
		panic(err)
	}
}

func writeMissingParamError(w http.ResponseWriter, paramName string) {
	http.Error(w, fmt.Sprintf("missing query param %q", paramName), http.StatusBadRequest)
}

func errToStatusCode(err error) int {
	// switch err {
	// case xorm.ErrRecordNotFound:
	// 	return http.StatusNotFound
	// default:
	return http.StatusInternalServerError
	// }
}
