package entity

import "time"

type Documentation struct {
	name        string
	date        *time.Time
	attendant   uint32
	location    string
	description string
	participant uint32
	photo       *Photo
	organizer   *Branch
}

type Photo struct {
	url  string
	name string
	id   string
}
