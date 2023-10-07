package service

import (
	"context"
	"project-adhyaksa/pkg/pagination"
	queryfilter "project-adhyaksa/services/event/domain/query_filter"
	"project-adhyaksa/services/event/domain/repository"
	"project-adhyaksa/services/event/domain/service"
	"project-adhyaksa/services/event/internal/service/mapping"
)

type eventService struct {
	eventRepository repository.EventRepository
}

func NewEventService(eventRepository repository.EventRepository) service.EventService {
	return &eventService{eventRepository: eventRepository}
}

func (uc *eventService) Create(event service.EventServiceDTO, ctx context.Context) error {

	eventEntity, err := mapping.EventMappingServiceDTOEntity(&event)
	if err != nil {
		return err
	}

	return uc.eventRepository.Create(*eventEntity, ctx)
}

func (uc *eventService) GetListPaginated(ctx context.Context,
	pagin *pagination.Paginator,
	filter *queryfilter.GetEventQueryFilter,
) ([]service.EventServiceDTO, error) {

	eventEntities, err := uc.eventRepository.GetListPaginated(ctx, pagin, filter)
	if err != nil {
		return nil, err
	}

	result := mapping.EventMappingEntityServiceDTOList(&eventEntities)

	return result, nil

}
