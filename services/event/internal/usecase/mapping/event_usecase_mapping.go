package mapping

import (
	"project-adhyaksa/services/event/domain/service"
	"project-adhyaksa/services/event/domain/usecase"
)

func EventMappingUsecaseToService(event *usecase.EventUseCaseDTO) service.EventServiceDTO {
	return service.EventServiceDTO{
		BranchID:    event.BranchID,
		AdminID:     event.AdminID,
		Name:        event.Name,
		StartTime:   event.StartTime,
		EndTime:     event.EndTime,
		Location:    event.Location,
		Description: event.Description,
	}
}

func EventMappingServiceToUsecaseList(events *[]service.EventServiceDTO) []usecase.EventUseCaseDTO {
	var eventUsecases = make([]usecase.EventUseCaseDTO, len(*events))
	for i, event := range *events {
		eventusecase := usecase.EventUseCaseDTO{
			ID:          event.ID,
			BranchID:    event.BranchID,
			BranchName:  event.BranchName,
			AdminID:     event.AdminID,
			Name:        event.Name,
			StartTime:   event.StartTime,
			EndTime:     event.EndTime,
			Location:    event.Location,
			Description: event.Description,
		}
		eventUsecases[i] = eventusecase
	}
	return eventUsecases
}

func EventMappingServiceToUsecase(event *service.EventServiceDTO) *usecase.EventUseCaseDTO {
	eventUsecase := usecase.EventUseCaseDTO{
		ID:          event.ID,
		BranchID:    event.BranchID,
		BranchName:  event.BranchName,
		AdminID:     event.AdminID,
		Name:        event.Name,
		StartTime:   event.StartTime,
		EndTime:     event.EndTime,
		Location:    event.Location,
		Description: event.Description,
	}

	return &eventUsecase
}
