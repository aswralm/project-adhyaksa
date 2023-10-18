package service

import (
	"context"
	"mime/multipart"
	"project-adhyaksa/pkg/pagination"
	"time"
)

type DocumentationServiceDTO struct {
	ID          string
	BranchID    string
	AdminID     string
	Name        string
	Date        *time.Time
	Location    string
	Description string
	Participant uint32

	PhotoService []*PhotoServiceDTO
}

type DocumentationService interface {
	Create(documentation DocumentationServiceDTO, file multipart.File, ctx context.Context) error

	GetListPaginated(
		pagin *pagination.Paginator,
		ctx context.Context,
	) ([]*DocumentationServiceDTO, error)

	GetByID(id string, ctx context.Context) (*DocumentationServiceDTO, error)
}
