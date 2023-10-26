package repository

import (
	"project-adhyaksa/pkg/config"
	"project-adhyaksa/services/event/domain/repository"
)

type Repository struct {
	Config                  *config.Config
	EventRepository         repository.EventRepository
	DocumentationRepository repository.DocumentationRepository
	ParticipantRepository   repository.ParticipantRepository
}

func InitRepository(config *config.Config) *Repository {
	eventRepository := NewEventRepository(config)
	documentationRepository := NewDocumentationRepository(config)
	participantRepository := NewParticipantRepository(config)

	return &Repository{
		Config:                  config,
		EventRepository:         eventRepository,
		DocumentationRepository: documentationRepository,
		ParticipantRepository:   participantRepository,
	}
}
