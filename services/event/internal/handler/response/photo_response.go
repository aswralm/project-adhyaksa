package response

import (
	"project-adhyaksa/services/event/domain/usecase"
)

type PhotoResponse struct {
	PhotoID       string `json:"photo_id"`
	PhotoPublicID string `json:"photo_public_id"`
	PhotoURL      string `json:"photo_url"`
	PhotoName     string `json:"photo_name"`
}

func ListPhoto(photoUseCases []*usecase.PhotoUseCaseDTO) []*PhotoResponse {
	var (
		photoResponses = make([]*PhotoResponse, len(photoUseCases))
	)
	for i, photoUseCase := range photoUseCases {
		photoResponse := DetailPhoto(photoUseCase)
		photoResponses[i] = &photoResponse
	}

	return photoResponses
}

func DetailPhoto(photoUsecase *usecase.PhotoUseCaseDTO) PhotoResponse {
	photoResponse := PhotoResponse{
		PhotoID:       photoUsecase.PhotoID,
		PhotoPublicID: photoUsecase.PhotoPublicID,
		PhotoURL:      photoUsecase.PhotoURL,
		PhotoName:     photoUsecase.PhotoName,
	}
	return photoResponse
}
