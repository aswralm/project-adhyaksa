package service

import (
	"project-adhyaksa/services/event/domain/service"
	"project-adhyaksa/services/event/internal/repository"
)

type Service struct {
	EventService service.EventService
}

func InitService(repository *repository.Repository) *Service {
	eventService := NewEventService(repository.EventRepository)

	return &Service{
		EventService: eventService,
	}
}
