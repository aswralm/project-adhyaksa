package mapping

import (
	createid "project-adhyaksa/pkg/create-id"
	"project-adhyaksa/services/event/domain/entity"
	"project-adhyaksa/services/event/domain/service"
)

func ParticipantMappingServiceDTOEntity(dto *service.ParticipantServiceDTO) (*entity.Participant, error) {

	eventEntity := entity.Event{}
	eventEntity.SetID(dto.EventID)

	participantEntity, err := entity.NewParticipant(entity.ParticipantDTO{
		ID:     createid.CreateID(),
		UserID: dto.UserID,
		Status: dto.Status,
		Event:  &eventEntity,
	})
	if err != nil {
		return nil, err
	}

	return participantEntity, nil
}
