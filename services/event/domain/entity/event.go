package entity

import (
	"errors"
	"project-adhyaksa/services/event/internal/customerror"
	"time"

	"github.com/google/uuid"
)

type Event struct {
	id          string
	adminID     string
	name        string
	startTime   *time.Time
	endTime     *time.Time
	location    string
	description string

	//relations
	branch *Branch
}

type EventDTO struct {
	ID          string
	Name        string
	StartTime   *time.Time
	EndTime     *time.Time
	Location    string
	Description string

	Organizer *Branch
}

// mapping for DTO to Entity
func NewEvent(event EventDTO) (*Event, error) {

	if event.Name == "" || event.StartTime == nil || event.EndTime == nil || event.Location == "" || event.Description == "" {
		return nil, errors.New(customerror.ERROR_FIELD_ENTITY)
	}

	return &Event{
		id:          event.ID,
		name:        event.Name,
		startTime:   event.StartTime,
		endTime:     event.EndTime,
		location:    event.Location,
		description: event.Description,
		branch:      event.Organizer,
	}, nil
}

// getter & setter for entity
func (e Event) SetID() {
	e.id = uuid.New().String()
}

func (e *Event) GetID() string {
	return e.id
}

func (e Event) SetName(name string) {
	e.name = name
}

func (e *Event) GetName() string {
	return e.name
}

func (e Event) SetStartTime(startTime *time.Time) {
	e.startTime = startTime
}

func (e *Event) GetStartTime() *time.Time {
	return e.startTime
}

func (e Event) SetEndTime(endTime *time.Time) {
	e.endTime = endTime
}

func (e *Event) GetEndTime() *time.Time {
	return e.endTime
}

func (e Event) SetLocation(location string) {
	e.location = location
}

func (e *Event) GetLocation() string {
	return e.location
}

func (e Event) SetDescription(description string) {
	e.description = description
}

func (e *Event) GetDescription() string {
	return e.description
}

func (e *Event) GetAdminID() string {
	return e.adminID
}

func (e *Event) SetAdminID(adminID string) {
	e.adminID = adminID
}

func (e *Event) GetBranch() *Branch {
	return e.branch
}

func (e *Event) SetBranch(branch Branch) {
	e.branch = &branch
}
