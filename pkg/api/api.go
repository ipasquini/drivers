package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/ipasquini/drivers/pkg/data"
	"github.com/ipasquini/drivers/pkg/database"
	scribble "github.com/nanobox-io/golang-scribble"
	"io/ioutil"
	"net/http"
)

type api struct {
	router   http.Handler
	database *database.Database
}

type API interface {
	Router() http.Handler
}

func New(scribble *scribble.Driver) API {
	api := &api{database: &database.Database{Scribble: scribble}}

	router := mux.NewRouter()
	router.HandleFunc("/drivers/{ID:[0-9]+}", api.get).Methods(http.MethodGet)
	router.HandleFunc("/drivers", api.post).Methods(http.MethodPost)
	api.router = router

	return api
}

func (api *api) Router() http.Handler {
	return api.router
}

func (api *api) get(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	channel := make(chan *data.DriverWithError)
	go api.database.Read(mux.Vars(request)["ID"], channel)

	driverWithError := <-channel
	if driverWithError.Err != nil {
		writer.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(writer).Encode("Driver not found")
		return
	}

	writer.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(writer).Encode(driverWithError.Driver)
}

func (api *api) post(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(writer).Encode("Error getting request body")
		return
	}

	driver := &data.Driver{}
	err = json.Unmarshal(body, &driver)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(writer).Encode("Invalid driver")
		return
	}

	go api.database.Write(driver)
	writer.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(writer).Encode(driver)
}
