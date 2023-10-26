package usecase

import (
	"context"
	"project-adhyaksa/services/event/domain/service"
	"project-adhyaksa/services/event/domain/usecase"
	"project-adhyaksa/services/event/internal/usecase/mapping"
)

type participantUseCase struct {
	participantService service.ParticipantService
	evenService        service.EventService
}

func NewParticipantUseCase(participantService service.ParticipantService, evenService service.EventService) usecase.ParticipantUseCase {
	return &participantUseCase{participantService: participantService, evenService: evenService}
}

func (u *participantUseCase) ConfirmEvent(participant usecase.ParticipantUseCaseDTO, ctx context.Context) error {
	participantService := mapping.ParticipantMappingUsecaseToService(&participant)

	_, err := u.evenService.GetByID(participant.EventID, ctx)
	if err != nil {
		return err
	}

	return u.participantService.Create(participantService, ctx)
}
