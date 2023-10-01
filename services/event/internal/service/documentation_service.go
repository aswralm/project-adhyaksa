package service

import (
	"context"
	"mime/multipart"
	createid "project-adhyaksa/pkg/create-id"
	"project-adhyaksa/pkg/upload"
	"project-adhyaksa/services/event/domain/entity"
	"project-adhyaksa/services/event/domain/repository"
	"project-adhyaksa/services/event/domain/service"
	"project-adhyaksa/services/event/internal/customerror"
)

type documentationService struct {
	documentationRepository repository.DocumentationRepository
	upload                  upload.Upload
}

func NewDocumentationService(documentationRepository repository.DocumentationRepository, upload upload.Upload) service.DocumentatitonService {
	return &documentationService{documentationRepository: documentationRepository, upload: upload}
}

func (s *documentationService) Create(documentation service.DocumentationServiceDTO, file *multipart.File, ctx context.Context) error {
	documentationEntity, err := entity.NewDocumentation(entity.DocumentationDTO{
		ID:          createid.CreateID(),
		AdminID:     documentation.AdminID,
		Name:        documentation.Name,
		Date:        documentation.Date,
		Location:    documentation.Location,
		Description: documentation.Description,
		Participant: documentation.Participant,
	})
	if err != nil {
		return &customerror.Err{
			Code:   customerror.ERROR_INVALID_REQUEST,
			Errors: err.Error(),
		}
	}

	photoEntity, err := entity.NewPhoto(entity.PhotoDTO{
		ID:            createid.CreateID(),
		PublicID:      documentation.PhotoPublicID,
		URL:           documentation.PhotoURL,
		Name:          documentation.PhotoName,
		Documentation: documentationEntity,
	})
	if err != nil {
		return &customerror.Err{
			Code:   customerror.ERROR_INVALID_REQUEST,
			Errors: err.Error(),
		}
	}

	s.upload.UploadImage(ctx, file)
	if err := s.documentationRepository.Create(*photoEntity, ctx); err != nil {
		s.upload.RemoveImage(ctx, photoEntity.GetPublicID())
		return err
	}

	return nil
}
