package main

import (
	"github.com/ipasquini/drivers/pkg/server"
	"github.com/nanobox-io/golang-scribble"
	"log"
	"net/http"
)

func main() {
	database, _ := scribble.New("./driver", nil)
	api := server.New(database)
	log.Fatal(http.ListenAndServe(":8080", api.Router()))
}
