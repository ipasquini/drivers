package service

import (
	"github.com/ipasquini/drivers/pkg/data"
	scribble "github.com/nanobox-io/golang-scribble"
	"strconv"
)

type Database struct {
	Scribble *scribble.Driver
}

func (database *Database) Read(id string) (*data.Driver, error) {
	driver := &data.Driver{}
	err := database.Scribble.Read("driver", id, &driver)
	return driver, err
}

func (database *Database) Write(driver *data.Driver) error {
	return database.Scribble.Write("driver", strconv.Itoa(driver.ID), driver)
}
