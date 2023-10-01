package repository

import (
	"context"
	"project-adhyaksa/services/event/domain/entity"
)

type EventRepository interface {
	Create(event entity.Event, ctx context.Context) error
}
