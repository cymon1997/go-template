package db

import (
	"log"
	"sync"
	"time"
)

var (
	loc     *time.Location
	syncLoc sync.Once

	NowUTC = func() time.Time {
		return time.Now().UTC()
	}

	NowTZ = func() time.Time {
		return time.Now().In(loc)
	}
)

func InitTimezone(tz string) {
	syncLoc.Do(func() {
		var err error
		loc, err = time.LoadLocation(tz)
		if err != nil {
			log.Fatal("error load db timezone: ", err)
		}
	})
}
