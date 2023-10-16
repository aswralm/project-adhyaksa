package usecase

import (
	"context"
	"project-adhyaksa/pkg/pagination"
	queryfilter "project-adhyaksa/services/event/domain/query_filter"
	"project-adhyaksa/services/event/domain/service"
	"project-adhyaksa/services/event/domain/usecase"
	"project-adhyaksa/services/event/internal/usecase/mapping"
)

type eventUseCase struct {
	eventService service.EventService
}

func NewEventUseCase(eventService service.EventService) usecase.EventUseCase {
	return &eventUseCase{eventService: eventService}
}

func (uc *eventUseCase) Create(event usecase.EventUseCaseDTO, ctx context.Context) error {
	eventDTO := mapping.EventMappingUsecaseToService(&event)
	return uc.eventService.Create(eventDTO, ctx)
}

func (uc *eventUseCase) GetListPaginated(
	pagin *pagination.Paginator,
	filter *queryfilter.GetEventQueryFilter,
) ([]usecase.EventUseCaseDTO, error) {
	eventService, err := uc.eventService.GetListPaginated(pagin, filter)
	if err != nil {
		return nil, err
	}
	result := mapping.EventMappingServiceToUsecaseList(&eventService)

	return result, nil
}

func (uc *eventUseCase) GetByID(id string, ctx context.Context) (*usecase.EventUseCaseDTO, error) {
	eventServiceDTO, err := uc.eventService.GetByID(id, ctx)
	if err != nil {
		return nil, err
	}
	result := mapping.EventMappingServiceToUsecase(eventServiceDTO)
	return result, nil

}
