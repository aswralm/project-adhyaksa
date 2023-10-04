package repository

import (
	"context"
	"project-adhyaksa/services/event/domain/entity"
)

type DocumentationRepository interface {
	Create(documentation entity.Documentation, photo entity.Photo, ctx context.Context) error
}
