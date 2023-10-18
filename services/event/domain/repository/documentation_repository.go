package repository

import (
	"context"
	"project-adhyaksa/pkg/pagination"
	"project-adhyaksa/services/event/domain/entity"
)

type DocumentationRepository interface {
	Create(documentation entity.Documentation, photo entity.Photo, ctx context.Context) error

	GetListPaginated(
		pagin *pagination.Paginator,
		ctx context.Context,
	) ([]*entity.Documentation, error)

	GetByID(id string, ctx context.Context) (*entity.Documentation, error)
}
