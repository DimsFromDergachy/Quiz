package storage

import (
	"testing"
)

func TestGetEvents(*testing.T) {
	connect()

	createTables()

	//insertEvent(true)
	//insertEvent(false)

	getEvents()
}
