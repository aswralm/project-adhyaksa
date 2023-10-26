package service

import (
	"context"
)

type ParticipantServiceDTO struct {
	ID      string
	UserID  string
	EventID string
	Status  string
}
type ParticipantService interface {
	Create(participant ParticipantServiceDTO, ctx context.Context) error
}
