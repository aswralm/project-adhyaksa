package entity

import (
	"errors"
	"project-adhyaksa/services/event/internal/customerror"
	"time"

	"github.com/google/uuid"
)

type Documentation struct {
	id          string
	adminID     string
	name        string
	date        *time.Time
	location    string
	description string
	participant uint32

	//relations
	branch *Branch
	photos []*Photo
}
type DocumentationDTO struct {
	ID          string
	AdminID     string
	Name        string
	Date        *time.Time
	Location    string
	Description string
	Participant uint32

	//relation
	Photos []*Photo
	Branch *Branch
}

// mapping for DTO to Entity
func NewDocumentation(documentation DocumentationDTO) (*Documentation, error) {

	if documentation.Name == "" || documentation.Date == nil || documentation.Location == "" || documentation.Description == "" {
		return nil, &customerror.Err{
			Code:   customerror.ERROR_INVALID_REQUEST,
			Errors: errors.New(customerror.ERROR_FIELD_ENTITY).Error(),
		}
	}

	return &Documentation{
		id:          documentation.ID,
		adminID:     documentation.AdminID,
		name:        documentation.Name,
		date:        documentation.Date,
		location:    documentation.Location,
		description: documentation.Description,
		participant: documentation.Participant,
		branch:      documentation.Branch,
	}, nil
}

// getter & setter for entity
func (d *Documentation) SetID() {
	d.id = uuid.New().String()
}

func (d *Documentation) GetID() string {
	return d.id
}

func (d *Documentation) SetAdminID(adminID string) {
	d.adminID = adminID
}

func (d *Documentation) GetAdminID() string {
	return d.adminID
}

func (d *Documentation) SetName(name string) {
	d.name = name
}

func (d *Documentation) GetName() string {
	return d.name
}

func (d *Documentation) SetDate(date *time.Time) {
	d.date = date
}

func (d *Documentation) GetDate() *time.Time {
	return d.date
}

func (d *Documentation) SetLocation(location string) {
	d.location = location
}

func (d *Documentation) GetLocation() string {
	return d.location
}

func (d *Documentation) SetDescription(description string) {
	d.description = description
}

func (d *Documentation) GetDescription() string {
	return d.description
}

func (d *Documentation) SetParticipant(participant uint32) {
	d.participant = participant
}

func (d *Documentation) GetParticipant() uint32 {
	return d.participant
}

func (e *Documentation) GetBranch() *Branch {
	return e.branch
}

func (e *Documentation) SetBranch(branch Branch) {
	e.branch = &branch
}

func (e *Documentation) GetPhoto() []*Photo {
	return e.photos
}

func (e *Documentation) SetPhoto(photo []*Photo) {
	e.photos = photo
}
