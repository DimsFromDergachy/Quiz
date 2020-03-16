package storage

import (
	"testing"
	"time"
)

func TestGetEvents(*testing.T) {
	connect()

	//dropTable()
	//createTables()

	insertEvent(event{
		title:  "Quiz I",
		active: true,
		date:   time.Now(),
	})
	insertEvent(event{
		title:  "Quiz II",
		active: false,
		date:   time.Now().Add((-1) * 24 * time.Hour),
	})

	getEvents()
}
