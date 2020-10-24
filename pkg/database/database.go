package database

import (
	"github.com/ipasquini/drivers/pkg/data"
	"github.com/jpillora/backoff"
	scribble "github.com/nanobox-io/golang-scribble"
	"strconv"
	"time"
)

type Database struct {
	Scribble *scribble.Driver
}

func (database *Database) Read(id string, channel chan *data.DriverWithError) {
	driver := &data.Driver{}
	err := database.Scribble.Read("driver", id, &driver)
	channel <- &data.DriverWithError{Driver: driver, Err: err}
}

func (database *Database) Write(driver *data.Driver) {
	exponentialBackoff := &backoff.Backoff{
		Min:    100 * time.Millisecond,
		Max:    12 * time.Hour,
		Factor: 2,
		Jitter: true,
	}

	timeToWait := exponentialBackoff.Duration()
	for ok := true; ok; ok = timeToWait.Hours() != 12.0 {
		timeToWait = exponentialBackoff.Duration()
		err := database.Scribble.Write("driver", strconv.Itoa(driver.ID), driver)
		if err == nil {
			break
		}
	}
}
