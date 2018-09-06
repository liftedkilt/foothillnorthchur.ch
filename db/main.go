package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/liftedkilt/foothillnorthchur.ch/db/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/julienschmidt/httprouter"
)

const (
	port = ":8888"
)

var (
	addr = flag.String("addr", "lyrics:lyrics@tcp(localhost:3306)/lyrics?charset=utf8", "the address of the database")
)

func main() {
	flag.Parse()

	db := setupDB(*addr)
	defer db.Close()

	router := httprouter.New()

	server := NewServer(db)
	server.RegisterRouter(router)

	log.Fatal(http.ListenAndServe(port, router))
}

func setupDB(addr string) *xorm.Engine {
	db, err := xorm.NewEngine("mysql", addr)
	if err != nil {
		panic(fmt.Sprintf("failed to connect to database: %v", err))
	}

	// db.CreateTables(model.Songs{})
	db.Sync(new(model.Songs))
	db.Sync(new(model.Playlist))

	return db
}
