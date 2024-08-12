package models

import "time"

type Event struct {
	//Model
	Id          uint
	Name        string
	TypeId      uint
	LocationId  uint
	Capacity    uint
	Description string
	PerformerId uint
	Duration    time.Duration
	StartTime   time.Time
	EndTime     time.Time
}
