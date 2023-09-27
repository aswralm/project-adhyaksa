package usecase

import (
	"project-adhyaksa/services/event/domain/usecase"
	"project-adhyaksa/services/event/internal/service"
)

type UseCase struct {
	EventUseCase usecase.EventUseCase
}

func InitUseCase(service *service.Service) *UseCase {
	eventUseCase := NewEventUseCase(service.EventService)

	return &UseCase{
		EventUseCase: eventUseCase,
	}
}
