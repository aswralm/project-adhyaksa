package response

import "project-adhyaksa/services/event/domain/usecase"

type EventResponse struct {
	ID          string `json:"id"`
	BranchID    string `json:"branch_id"`
	BranchName  string `json:"branch_name"`
	AdminID     string `json:"admin_id"`
	Name        string `json:"event_name"`
	StartTime   string `json:"start_event"`
	EndTime     string `json:"end_event"`
	Location    string `json:"location"`
	Description string `json:"description"`
}

func ListMapping(events *[]usecase.EventUseCaseDTO) []EventResponse {
	var eventReponses = make([]EventResponse, len(*events))

	for i, event := range *events {
		eventReponse := EventResponse{
			ID:          event.ID,
			BranchID:    event.BranchID,
			BranchName:  event.BranchName,
			AdminID:     event.AdminID,
			Name:        event.Name,
			StartTime:   event.StartTime.String(),
			EndTime:     event.EndTime.String(),
			Location:    event.Location,
			Description: event.Description,
		}
		eventReponses[i] = eventReponse
	}
	return eventReponses
}
