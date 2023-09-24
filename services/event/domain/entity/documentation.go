package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Documentation struct {
	id          string
	adminID     string
	name        string
	date        *time.Time
	attendant   uint32
	location    string
	description string
	participant uint32
}
type DocumentationDTO struct {
	ID          string
	AdminID     string
	Name        string
	Date        *time.Time
	Attendant   uint32
	Location    string
	Description string
	Participant uint32

	//relation
	photo  *[]Photo
	branch *Branch
}

// mapping for DTO to Entity
func NewDocumentation(documentation DocumentationDTO) (*Documentation, error) {
	if documentation.ID == "" || documentation.AdminID == "" || documentation.Name == "" || documentation.Date == nil || documentation.Location == "" || documentation.Description == "" {
		return nil, errors.New("ERROR_FIELD_ENTITY")
	}

	return &Documentation{
		id:          documentation.ID,
		adminID:     documentation.AdminID,
		name:        documentation.Name,
		date:        documentation.Date,
		attendant:   documentation.Attendant,
		location:    documentation.Location,
		description: documentation.Description,
		participant: documentation.Participant,
	}, nil
}

// getter & setter for entity
func (d Documentation) SetID() {
	d.id = uuid.New().String()
}

func (d *Documentation) GetID() string {
	return d.id
}

func (d Documentation) SetAdminID(adminID string) {
	d.adminID = adminID
}

func (d *Documentation) GetAdminID() string {
	return d.adminID
}

func (d Documentation) SetName(name string) {
	d.name = name
}

func (d *Documentation) GetName() string {
	return d.name
}

func (d Documentation) SetDate(date *time.Time) {
	d.date = date
}

func (d *Documentation) GetDate() *time.Time {
	return d.date
}

func (d Documentation) SetAttendant(attendant uint32) {
	d.attendant = attendant
}

func (d *Documentation) GetAttendant() uint32 {
	return d.attendant
}

func (d Documentation) SetLocation(location string) {
	d.location = location
}

func (d *Documentation) GetLocation() string {
	return d.location
}

func (d Documentation) SetDescription(description string) {
	d.description = description
}

func (d *Documentation) GetDescription() string {
	return d.description
}

func (d Documentation) SetParticipant(participant uint32) {
	d.participant = participant
}

func (d *Documentation) GetParticipant() uint32 {
	return d.participant
}
