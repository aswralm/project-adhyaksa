package service

import (
	"context"
	"fmt"
	"mime/multipart"
	createid "project-adhyaksa/pkg/create-id"
	"project-adhyaksa/pkg/pagination"
	"project-adhyaksa/pkg/upload"
	"project-adhyaksa/services/event/domain/entity"
	"project-adhyaksa/services/event/domain/repository"
	"project-adhyaksa/services/event/domain/service"
	"project-adhyaksa/services/event/internal/service/mapping"
)

type documentationService struct {
	documentationRepository repository.DocumentationRepository
	upload                  upload.Upload
}

func NewDocumentationService(documentationRepository repository.DocumentationRepository, upload upload.Upload) service.DocumentationService {
	return &documentationService{documentationRepository: documentationRepository, upload: upload}
}

func (s *documentationService) Create(documentation service.DocumentationServiceDTO, file multipart.File, ctx context.Context) error {
	fmt.Println(documentation)
	branch, err := entity.NewBranch(entity.BranchDTO{ID: documentation.BranchID})
	if err != nil {
		return err
	}
	documentationEntity, err := entity.NewDocumentation(entity.DocumentationDTO{
		ID:          createid.CreateID(),
		AdminID:     documentation.AdminID,
		Name:        documentation.Name,
		Date:        documentation.Date,
		Location:    documentation.Location,
		Description: documentation.Description,
		Participant: documentation.Participant,
		Branch:      branch,
	})

	if err != nil {
		return err
	}

	photoEntity, err := entity.NewPhoto(entity.PhotoDTO{
		ID:            createid.CreateID(),
		Documentation: documentationEntity,
	})
	if err != nil {
		return err
	}

	url, publicID, err := s.upload.UploadImage(ctx, file)
	if err != nil {
		return err
	}

	photoEntity.SetURL(string(url))
	photoEntity.SetPublicID(publicID)

	if err := s.documentationRepository.Create(*documentationEntity, *photoEntity, ctx); err != nil {
		if err := s.upload.RemoveImage(ctx, publicID); err != nil {
			return err
		}
		return err
	}

	return nil
}

func (s *documentationService) GetListPaginated(
	pagin *pagination.Paginator,
	ctx context.Context,
) ([]*service.DocumentationServiceDTO, error) {

	documentationEntities, err := s.documentationRepository.GetListPaginated(pagin, ctx)
	if err != nil {
		return nil, err
	}
	result := mapping.DocumentationMappingEntityServiceDTOList(documentationEntities)

	return result, nil
}

func (s *documentationService) GetByID(id string, ctx context.Context) (*service.DocumentationServiceDTO, error) {
	documentationRepository, err := s.documentationRepository.GetByID(id, ctx)
	if err != nil {
		return nil, err
	}
	result := mapping.DocumentationMappingEntityServiceDTO(documentationRepository)

	return result, err
}
