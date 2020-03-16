package storage

import "time"

// Event is a structure which represente an ... event
type event struct {
	id     int
	title  string
	active bool
	date   time.Time
}
