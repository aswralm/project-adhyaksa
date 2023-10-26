package mapping

import (
	"project-adhyaksa/services/event/domain/service"
	"project-adhyaksa/services/event/domain/usecase"
)

func ParticipantMappingUsecaseToService(participant *usecase.ParticipantUseCaseDTO) service.ParticipantServiceDTO {
	return service.ParticipantServiceDTO{
		UserID:  participant.UserID,
		EventID: participant.EventID,
		Status:  participant.Status,
	}
}
