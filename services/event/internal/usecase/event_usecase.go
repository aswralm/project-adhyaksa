package usecase

import (
	"context"
	"project-adhyaksa/services/event/domain/service"
	"project-adhyaksa/services/event/domain/usecase"
)

type eventUseCase struct {
	eventService service.EventService
}

func NewEventUseCase(eventService service.EventService) usecase.EventUseCase {
	return &eventUseCase{eventService: eventService}
}

func (uc *eventUseCase) Create(event usecase.EventUseCaseDTO, ctx context.Context) error {
	eventDTO := service.EventServiceDTO{
		BranchID:    event.BranchID,
		AdminID:     event.AdminID,
		Name:        event.Name,
		StartTime:   event.StartTime,
		EndTime:     event.EndTime,
		Location:    event.Location,
		Description: event.Description,
	}
	return uc.eventService.Create(eventDTO, ctx)
}
