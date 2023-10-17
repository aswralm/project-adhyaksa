package mapping

import (
	"project-adhyaksa/services/event/domain/usecase"
	"project-adhyaksa/services/event/internal/handler/request"
)

func DocumentationRequestToUsecaseDTO(documentationRequest *request.RegisterDocumentationRequest, adminID string) usecase.DocumentationUseCaseDTO {
	documentation := usecase.DocumentationUseCaseDTO{
		BranchID:    documentationRequest.BranchID,
		AdminID:     adminID,
		Name:        documentationRequest.Name,
		Date:        documentationRequest.Date,
		Location:    documentationRequest.Location,
		Description: documentationRequest.Description,
		Participant: documentationRequest.Participant,
	}

	return documentation
}
