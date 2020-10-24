package main

import (
	"github.com/ipasquini/drivers/pkg/api"
	"github.com/nanobox-io/golang-scribble"
	"log"
	"net/http"
)

func main() {
	database, err := scribble.New("./driver", nil)
	if err == nil {
		log.Fatal(http.ListenAndServe(":8080", api.New(database).Router()))
	} else {
		log.Fatal("Unable to initialize database")
	}
}
