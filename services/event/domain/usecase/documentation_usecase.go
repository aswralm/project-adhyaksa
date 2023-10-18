package usecase

import (
	"context"
	"mime/multipart"
	"project-adhyaksa/pkg/pagination"
	"time"
)

type DocumentationUseCaseDTO struct {
	ID          string
	BranchID    string
	AdminID     string
	Name        string
	Date        *time.Time
	Location    string
	Description string
	Participant uint32

	Photos []*PhotoUseCaseDTO
}

type DocumentatitonUseCase interface {
	Create(documentation DocumentationUseCaseDTO, file multipart.File, ctx context.Context) error

	GetListPaginated(
		pagin *pagination.Paginator,
		ctx context.Context,
	) ([]*DocumentationUseCaseDTO, error)

	GetByID(id string, ctx context.Context) (*DocumentationUseCaseDTO, error)
}
