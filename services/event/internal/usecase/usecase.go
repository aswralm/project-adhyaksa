package usecase

import (
	"project-adhyaksa/services/event/domain/usecase"
	"project-adhyaksa/services/event/internal/service"
)

type UseCase struct {
	EventUseCase         usecase.EventUseCase
	DocumentationUseCase usecase.DocumentatitonUseCase
	ParticipantUseCase   usecase.ParticipantUseCase
}

func InitUseCase(service *service.Service) *UseCase {
	eventUseCase := NewEventUseCase(service.EventService)
	documentationUseCase := NewDocumentationUseCase(service.DocumentationService)
	participantUseCase := NewParticipantUseCase(service.ParticipantService, service.EventService)

	return &UseCase{
		EventUseCase:         eventUseCase,
		DocumentationUseCase: documentationUseCase,
		ParticipantUseCase:   participantUseCase,
	}
}
