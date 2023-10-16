package mapping

import (
	createid "project-adhyaksa/pkg/create-id"
	"project-adhyaksa/services/event/domain/entity"
	"project-adhyaksa/services/event/domain/service"
)

func EventMappingServiceDTOEntity(event *service.EventServiceDTO) (*entity.Event, error) {

	branch, err := entity.NewBranch(entity.BranchDTO{ID: event.BranchID})
	if err != nil {
		return nil, err
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
		return nil, err
	}

	return eventEntity, nil
}

func EventMappingEntityServiceDTOList(eventEntities []*entity.Event) []service.EventServiceDTO {
	var (
		eventServices = make([]service.EventServiceDTO, len(eventEntities))
	)
	for i, eventEntity := range eventEntities {
		eventService := service.EventServiceDTO{
			ID:          eventEntity.GetID(),
			BranchID:    eventEntity.GetBranch().GetID(),
			BranchName:  eventEntity.GetBranch().GetName(),
			AdminID:     eventEntity.GetAdminID(),
			Name:        eventEntity.GetName(),
			StartTime:   eventEntity.GetStartTime(),
			EndTime:     eventEntity.GetEndTime(),
			Location:    eventEntity.GetLocation(),
			Description: eventEntity.GetDescription(),
		}
		eventServices[i] = eventService
	}

	return eventServices
}

func EventMappingEntityServiceDTO(eventEntity *entity.Event) *service.EventServiceDTO {

	eventService := service.EventServiceDTO{
		ID:          eventEntity.GetID(),
		BranchID:    eventEntity.GetBranch().GetID(),
		BranchName:  eventEntity.GetBranch().GetName(),
		AdminID:     eventEntity.GetAdminID(),
		Name:        eventEntity.GetName(),
		StartTime:   eventEntity.GetStartTime(),
		EndTime:     eventEntity.GetEndTime(),
		Location:    eventEntity.GetLocation(),
		Description: eventEntity.GetDescription(),
	}

	return &eventService
}
