package mapping

import (
	"project-adhyaksa/services/event/domain/entity"
	"project-adhyaksa/services/event/domain/service"
)

func DocumentationMappingEntityServiceDTOList(documentationEntities []*entity.Documentation) []*service.DocumentationServiceDTO {
	var (
		documentationServices = make([]*service.DocumentationServiceDTO, len(documentationEntities))
	)
	for i, documentationEntity := range documentationEntities {
		documentationService := DocumentationMappingEntityServiceDTO(documentationEntity)
		documentationServices[i] = documentationService
	}

	return documentationServices
}

func DocumentationMappingEntityServiceDTO(documenationEntity *entity.Documentation) *service.DocumentationServiceDTO {
	if documenationEntity == nil {
		return nil
	}
	eventService := service.DocumentationServiceDTO{
		ID:           documenationEntity.GetID(),
		BranchID:     documenationEntity.GetBranch().GetID(),
		AdminID:      documenationEntity.GetAdminID(),
		PhotoService: PhotoMappingEntityServiceDTOList(documenationEntity.GetPhoto()),
		Name:         documenationEntity.GetName(),
		Date:         documenationEntity.GetDate(),
		Location:     documenationEntity.GetLocation(),
		Description:  documenationEntity.GetDescription(),
	}

	return &eventService
}
