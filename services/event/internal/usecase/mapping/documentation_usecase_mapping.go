package mapping

import (
	"project-adhyaksa/services/event/domain/service"
	"project-adhyaksa/services/event/domain/usecase"
)

func DocumentationMappingServiceToUseCaseList(documentationServices *[]service.DocumentationServiceDTO) *[]usecase.DocumentationUseCaseDTO {
	var (
		documentationUseCases = make([]usecase.DocumentationUseCaseDTO, len(*documentationServices))
	)
	for i, documentationService := range *documentationServices {
		documentationUseCase := DocumentationMappingServiceToUseCase(&documentationService)
		documentationUseCases[i] = *documentationUseCase
	}
	return &documentationUseCases
}

func DocumentationMappingServiceToUseCase(documentationService *service.DocumentationServiceDTO) *usecase.DocumentationUseCaseDTO {

	documentationUseCase := usecase.DocumentationUseCaseDTO{
		ID:          documentationService.ID,
		BranchID:    documentationService.BranchID,
		AdminID:     documentationService.AdminID,
		Photos:      PhotoMappingServiceToUseCaseList(documentationService.PhotoService),
		Name:        documentationService.Name,
		Date:        documentationService.Date,
		Location:    documentationService.Location,
		Description: documentationService.Description,
	}

	return &documentationUseCase
}
