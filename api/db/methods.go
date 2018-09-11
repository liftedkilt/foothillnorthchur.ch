package db

import (
	cache "github.com/patrickmn/go-cache"
)

func CreatePlaylist(date string, playlist Playlist) {
	store.Set(date, playlist, cache.DefaultExpiration)
	return
}

func GetPlaylist(date string) (playlist Playlist) {
	result, ok := store.Get(date)
	if ok {
		return result.(Playlist)
	} else {
		return
	}
}
