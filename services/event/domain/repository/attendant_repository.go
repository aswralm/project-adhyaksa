package repository

import (
	"context"
	"project-adhyaksa/services/event/domain/entity"
)

type ParticipantRepository interface {
	Create(participant *entity.Participant, ctx context.Context) error
}
