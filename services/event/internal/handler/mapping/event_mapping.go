package mapping

import (
	"project-adhyaksa/services/event/domain/usecase"
	"project-adhyaksa/services/event/internal/handler/request"
)

func EventRequestToUsecaseDTO(eventRequest *request.RegisterEventRequest) usecase.EventUseCaseDTO {
	return usecase.EventUseCaseDTO{
		BranchID:    eventRequest.BranchID,
		Name:        eventRequest.Name,
		StartTime:   eventRequest.StartTime,
		EndTime:     eventRequest.EndTime,
		Location:    eventRequest.Location,
		Description: eventRequest.Description,
	}
}
