package types

import "time"

type Event struct {
	id int
	time time.Time
	eventType string
}
