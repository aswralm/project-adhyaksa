package service

import (
	"context"
	createid "project-adhyaksa/pkg/create-id"
	"project-adhyaksa/services/event/domain/repository"
	"project-adhyaksa/services/event/domain/service"
	"project-adhyaksa/services/event/internal/service/mapping"
)

type participantService struct {
	participantRepository repository.ParticipantRepository
}

func NewParticipantService(participantRepository repository.ParticipantRepository) service.ParticipantService {
	return &participantService{participantRepository: participantRepository}
}

func (s *participantService) Create(participant service.ParticipantServiceDTO, ctx context.Context) error {
	entity, err := mapping.ParticipantMappingServiceDTOEntity(&participant)
	if err != nil {
		return err
	}
	entity.SetID(createid.CreateID())

	if err := s.participantRepository.Create(entity, ctx); err != nil {
		return err
	}
	return nil
}
