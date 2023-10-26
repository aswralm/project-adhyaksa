package usecase

import (
	"context"
	"project-adhyaksa/services/event/domain/service"
	"project-adhyaksa/services/event/domain/usecase"
	"project-adhyaksa/services/event/internal/usecase/mapping"
)

type participantUseCase struct {
	participantService service.ParticipantService
}

func NewParticipantUseCase(participantService service.ParticipantService) usecase.ParticipantUseCase {
	return &participantUseCase{participantService: participantService}
}

func (u *participantUseCase) ConfirmEvent(participant usecase.ParticipantUseCaseDTO, ctx context.Context) error {
	participantService := mapping.ParticipantMappingUsecaseToService(&participant)

	return u.participantService.Create(participantService, ctx)
}
