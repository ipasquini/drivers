package service

import (
	"github.com/ipasquini/drivers/pkg/data"
	scribble "github.com/nanobox-io/golang-scribble"
	"strconv"
)

type Database struct {
	Scribble *scribble.Driver
}

func (database *Database) Read(id string, channel chan *data.DriverWithError) {
	driver := &data.Driver{}
	err := database.Scribble.Read("driver", id, &driver)
	channel <- &data.DriverWithError{Driver: driver, Err: err}
}

func (database *Database) Write(driver *data.Driver, channel chan error) {
	channel <- database.Scribble.Write("driver", strconv.Itoa(driver.ID), driver)
}
