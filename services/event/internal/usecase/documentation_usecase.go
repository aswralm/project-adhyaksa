package usecase

import (
	"context"
	"mime/multipart"
	"project-adhyaksa/services/event/domain/service"
	"project-adhyaksa/services/event/domain/usecase"
)

type documentationUseCase struct {
	documentationService service.DocumentationService
}

func NewDocumentationUseCase(documentationService service.DocumentationService) usecase.DocumentatitonUseCase {
	return &documentationUseCase{documentationService: documentationService}
}

func (uc *documentationUseCase) Create(documentation usecase.DocumentationUseCaseDTO, file multipart.File, ctx context.Context) error {
	data := service.DocumentationServiceDTO{
		BranchID:    documentation.BranchID,
		AdminID:     documentation.AdminID,
		Name:        documentation.Name,
		PhotoName:   documentation.PhotoName,
		Date:        documentation.Date,
		Location:    documentation.Location,
		Description: documentation.Description,
		Participant: documentation.Participant,
	}
	return uc.documentationService.Create(data, file, ctx)
}
