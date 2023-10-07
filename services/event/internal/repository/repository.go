package repository

import (
	"database/sql"
	"project-adhyaksa/pkg/config"
	"project-adhyaksa/services/event/domain/repository"
)

type Repository struct {
	Config                  *config.Config
	EventRepository         repository.EventRepository
	DocumentationRepository repository.DocumentationRepository
}

func InitRepository(db *sql.DB, config *config.Config) *Repository {
	eventRepository := NewEventRepository(db, config)
	documentationRepository := NewDocumentationRepository(db, config)

	return &Repository{
		Config:                  config,
		EventRepository:         eventRepository,
		DocumentationRepository: documentationRepository,
	}
}
