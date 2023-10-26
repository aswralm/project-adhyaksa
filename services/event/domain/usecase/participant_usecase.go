package usecase

import "context"

type ParticipantUseCaseDTO struct {
	ID      string
	UserID  string
	EventID string
	Status  string
}
type ParticipantUseCase interface {
	ConfirmEvent(participant ParticipantUseCaseDTO, ctx context.Context) error
}
