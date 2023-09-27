package service

import (
	"context"
	createid "project-adhyaksa/pkg/create-id"
	"project-adhyaksa/services/event/domain/entity"
	"project-adhyaksa/services/event/domain/repository"
	"project-adhyaksa/services/event/domain/service"
	"project-adhyaksa/services/event/internal/customerror"
)

type eventService struct {
	eventRepository repository.EventRepository
}

func NewEventService(eventRepository repository.EventRepository) service.EventService {
	return &eventService{eventRepository: eventRepository}
}

func (uc *eventService) Create(event service.EventServiceDTO, ctx context.Context) error {

	branch, err := entity.NewBranch(entity.BranchDTO{ID: event.BranchID})
	if err != nil {
		return &customerror.Err{
			Code:   customerror.ERROR_INVALID_REQUEST,
			Errors: err.Error(),
		}
	}

	eventEntity, err := entity.NewEvent(entity.EventDTO{
		ID:          createid.CreateID(),
		Name:        event.Name,
		StartTime:   event.StartTime,
		EndTime:     event.EndTime,
		Location:    event.Location,
		Description: event.Description,
		Organizer:   branch,
	})
	if err != nil {
		return &customerror.Err{
			Code:   customerror.ERROR_INVALID_REQUEST,
			Errors: err.Error(),
		}
	}

	return uc.eventRepository.Create(*eventEntity, ctx)
}
