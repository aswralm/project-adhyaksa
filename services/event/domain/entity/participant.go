package entity

import (
	"errors"
	"project-adhyaksa/services/event/internal/customerror"
)

type Participant struct {
	id      string
	userID  string
	adminID string
	status  string

	//relation
	event *Event
}
type ParticipantDTO struct {
	ID      string
	UserID  string
	AdminID string
	Status  string

	Event *Event
}

// mapping for DTO to Entity
func NewParticipant(participant ParticipantDTO) (*Participant, error) {
	if participant.ID == "" || participant.UserID == "" || participant.AdminID == "" || participant.Status == "" || participant.Event == nil {
		return nil, &customerror.Err{
			Code:   customerror.ERROR_INVALID_REQUEST,
			Errors: errors.New(customerror.ERROR_FIELD_ENTITY).Error(),
		}
	}

	return &Participant{
		id:      participant.ID,
		userID:  participant.UserID,
		adminID: participant.AdminID,
		status:  participant.Status,
		event:   participant.Event,
	}, nil
}

// getter & setter for entity
func (p *Participant) SetID(id string) {
	p.id = id
}

func (p *Participant) GetID() string {
	return p.id
}

func (p *Participant) SetUserID(userID string) {
	p.userID = userID
}

func (p *Participant) GetUserID() string {
	return p.userID
}

func (p *Participant) SetAdminID(adminID string) {
	p.adminID = adminID
}

func (p *Participant) GetAdminID() string {
	return p.adminID
}

func (p *Participant) SetStatus(status string) {
	p.status = status
}

func (p *Participant) GetStatus() string {
	return p.status
}

func (p *Participant) SetEvent(event *Event) {
	p.event = event
}

func (p *Participant) GetEvent() *Event {
	return p.event
}
