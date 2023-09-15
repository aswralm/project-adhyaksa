package entity

import "time"

type Documentation struct {
	id          string
	adminID     string
	name        string
	date        *time.Time
	attendant   uint32
	location    string
	description string
	participant uint32

	//relation
	photo  *[]Photo
	branch *Branch
}
