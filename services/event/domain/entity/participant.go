package entity

type Attendant struct {
	id      string
	userID  string
	adminID string
	status  string

	//relation
	event *Event
}
