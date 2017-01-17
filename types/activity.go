package types

import "time"

type Activity struct {
	id             int
	ActivityType   string
	startDateTime  time.Time
	stopDateTime   time.Time
	event          Event
	user           User
	ActivityConfig ActivityConfig
}

const (
	WORK = "work"
	EAT = "eat"
	RELAX = "relax"
	CUSTOM = "custom"
)

type ActivityConfig struct {
	id         int
	activityId int
	Duration   int64
	ChunkSize  int64
}
