package entity

import (
	"errors"
	valueobject "project-adhyaksa/services/event/domain/value_object"
	"project-adhyaksa/services/event/internal/customerror"
)

type Participant struct {
	id     string
	userID string
	status valueobject.StatusParticipant

	//relation
	event *Event
}
type ParticipantDTO struct {
	ID     string
	UserID string
	Status string

	Event *Event
}

// mapping for DTO to Entity
func NewParticipant(participant ParticipantDTO) (*Participant, error) {
	if participant.Status == "" {
		return nil, &customerror.Err{
			Code:   customerror.ERROR_INVALID_REQUEST,
			Errors: errors.New(customerror.ERROR_FIELD_ENTITY).Error(),
		}
	}

	result := &Participant{
		id:     participant.ID,
		userID: participant.UserID,
		event:  participant.Event,
	}

	if valueobject.IsValidStatusParticipant(valueobject.StatusParticipant(participant.Status)) {
		result.status = valueobject.StatusParticipant(participant.Status)
	}

	return result, nil
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

func (p *Participant) SetStatus(status valueobject.StatusParticipant) {
	p.status = status
}

func (p *Participant) GetStatus() valueobject.StatusParticipant {
	return p.status
}

func (p *Participant) SetEvent(event *Event) {
	p.event = event
}

func (p *Participant) GetEvent() *Event {
	return p.event
}
