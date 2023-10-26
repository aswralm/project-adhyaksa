package mapping

import (
	"project-adhyaksa/services/event/domain/usecase"
	"project-adhyaksa/services/event/internal/handler/request"
)

func ParticipantRequestToUsecaseDTO(participantRequest *request.ConfirmEventRequest) usecase.ParticipantUseCaseDTO {
	return usecase.ParticipantUseCaseDTO{
		EventID: participantRequest.EventID,
		Status:  participantRequest.Status,
	}
}
