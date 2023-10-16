package repository

import (
	"project-adhyaksa/pkg/config"
	"project-adhyaksa/services/event/domain/repository"
)

type Repository struct {
	Config                  *config.Config
	EventRepository         repository.EventRepository
	DocumentationRepository repository.DocumentationRepository
}

func InitRepository(config *config.Config) *Repository {
	eventRepository := NewEventRepository(config)
	documentationRepository := NewDocumentationRepository(config)

	return &Repository{
		Config:                  config,
		EventRepository:         eventRepository,
		DocumentationRepository: documentationRepository,
	}
}
