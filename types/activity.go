package types

import "time"

type Activity struct {
	id             int
	ActivityType   string
	StartDateTime  time.Time
	StopDateTime   time.Time
	event          Event
	User           User
	ActivityConfig ActivityConfig
	Launched       bool
}

const (
	WORK   = "work"
	EAT    = "eat"
	RELAX  = "relax"
	CUSTOM = "custom"
)

type ActivityConfig struct {
	id         int
	activityId int
	Duration   time.Duration
	ChunkSize  time.Duration
}
