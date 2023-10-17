package response

import "project-adhyaksa/services/event/domain/usecase"

type DocumentationResponse struct {
	ID          string           `json:"id"`
	BranchID    string           `json:"branch_id"`
	AdminID     string           `json:"admin_id"`
	Name        string           `json:"event_name"`
	Date        string           `json:"date"`
	Location    string           `json:"location"`
	Description string           `json:"description"`
	Photos      []*PhotoResponse `json:"photos"`
}

func ListDocumentation(documentationUseCases []*usecase.DocumentationUseCaseDTO) []DocumentationResponse {
	var documentationReponses = make([]DocumentationResponse, len(documentationUseCases))

	for i, documentationUseCase := range documentationUseCases {
		documentationReponse := DocumentationResponse{
			ID:          documentationUseCase.ID,
			BranchID:    documentationUseCase.BranchID,
			AdminID:     documentationUseCase.AdminID,
			Name:        documentationUseCase.Name,
			Date:        documentationUseCase.Date.String(),
			Location:    documentationUseCase.Location,
			Description: documentationUseCase.Description,
		}
		documentationReponses[i] = documentationReponse
	}
	return documentationReponses
}

func DetailDocumentation(documentationUseCase *usecase.DocumentationUseCaseDTO) *DocumentationResponse {

	photos := ListPhoto(documentationUseCase.Photos)
	documentationReponse := DocumentationResponse{
		ID:          documentationUseCase.ID,
		BranchID:    documentationUseCase.BranchID,
		AdminID:     documentationUseCase.AdminID,
		Name:        documentationUseCase.Name,
		Date:        documentationUseCase.Date.String(),
		Location:    documentationUseCase.Location,
		Description: documentationUseCase.Description,
		Photos:      photos,
	}

	return &documentationReponse
}
