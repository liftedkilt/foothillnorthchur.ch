package db

import (
	"time"

	cache "github.com/patrickmn/go-cache"
)

var store *cache.Cache

func init() {
	store = cache.New(8760*time.Hour, 10*time.Minute)
}
