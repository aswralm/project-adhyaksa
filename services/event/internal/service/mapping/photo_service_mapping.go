package mapping

import (
	"project-adhyaksa/services/event/domain/entity"
	"project-adhyaksa/services/event/domain/service"
)

func PhotoMappingEntityServiceDTOList(photoEntities *[]entity.Photo) *[]service.PhotoServiceDTO {
	var (
		photoServices = make([]service.PhotoServiceDTO, len(*photoEntities))
	)
	for i, photoEntity := range *photoEntities {
		photoService := PhotoMappingEntityServiceDTO(&photoEntity)
		photoServices[i] = *photoService
	}

	return &photoServices
}

func PhotoMappingEntityServiceDTO(photoEntity *entity.Photo) *service.PhotoServiceDTO {

	photoService := service.PhotoServiceDTO{
		PhotoID:       photoEntity.GetID(),
		PhotoPublicID: photoEntity.GetPublicID(),
		PhotoURL:      photoEntity.GetURL(),
		PhotoName:     photoEntity.GetName(),
	}

	return &photoService
}
