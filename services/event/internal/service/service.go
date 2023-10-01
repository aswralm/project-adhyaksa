package service

import (
	"project-adhyaksa/pkg/upload"
	"project-adhyaksa/services/event/domain/service"
	"project-adhyaksa/services/event/internal/repository"
)

type Service struct {
	EventService         service.EventService
	DocumentationService service.DocumentatitonService
}

func InitService(repository *repository.Repository, upload upload.Upload) *Service {
	eventService := NewEventService(repository.EventRepository)
	documentationService := NewDocumentationService(repository.DocumentationRepository, upload)

	return &Service{
		EventService:         eventService,
		DocumentationService: documentationService,
	}
}
