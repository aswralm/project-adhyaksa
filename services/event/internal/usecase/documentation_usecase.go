package usecase

import (
	"context"
	"mime/multipart"
	"project-adhyaksa/pkg/pagination"
	"project-adhyaksa/services/event/domain/service"
	"project-adhyaksa/services/event/domain/usecase"
	"project-adhyaksa/services/event/internal/usecase/mapping"
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
		Date:        documentation.Date,
		Location:    documentation.Location,
		Description: documentation.Description,
		Participant: documentation.Participant,
	}
	return uc.documentationService.Create(data, file, ctx)
}

func (uc *documentationUseCase) GetListPaginated(
	pagin *pagination.Paginator,
	ctx context.Context,
) (*[]usecase.DocumentationUseCaseDTO, error) {
	documentationServices, err := uc.documentationService.GetListPaginated(pagin, ctx)
	if err != nil {
		return nil, err
	}
	result := mapping.DocumentationMappingServiceToUseCaseList(documentationServices)

	return result, nil
}

func (uc *documentationUseCase) GetByID(id string, ctx context.Context) (*usecase.DocumentationUseCaseDTO, error) {
	documentationService, err := uc.documentationService.GetByID(id, ctx)
	if err != nil {
		return nil, err
	}
	result := mapping.DocumentationMappingServiceToUseCase(documentationService)

	return result, nil
}
