package mapping

import (
	"project-adhyaksa/services/event/domain/service"
	"project-adhyaksa/services/event/domain/usecase"
)

func PhotoMappingServiceToUseCaseList(photoServices []*service.PhotoServiceDTO) []*usecase.PhotoUseCaseDTO {
	var (
		photoUseCases = make([]*usecase.PhotoUseCaseDTO, len(photoServices))
	)
	for i, photoService := range photoServices {
		photoUseCase := PhotoMappingServiceToUseCase(photoService)
		photoUseCases[i] = photoUseCase
	}

	return photoUseCases
}

func PhotoMappingServiceToUseCase(photoService *service.PhotoServiceDTO) *usecase.PhotoUseCaseDTO {

	photoUseCase := usecase.PhotoUseCaseDTO{
		PhotoID:       photoService.PhotoID,
		PhotoPublicID: photoService.PhotoPublicID,
		PhotoURL:      photoService.PhotoURL,
		PhotoName:     photoService.PhotoName,
	}

	return &photoUseCase
}
