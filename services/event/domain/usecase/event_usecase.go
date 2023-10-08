package usecase

import (
	"context"
	"project-adhyaksa/pkg/pagination"
	queryfilter "project-adhyaksa/services/event/domain/query_filter"
	"time"
)

type EventUseCaseDTO struct {
	ID          string
	BranchID    string
	BranchName  string
	AdminID     string
	Name        string
	StartTime   *time.Time
	EndTime     *time.Time
	Location    string
	Description string
}

type EventUseCase interface {
	Create(event EventUseCaseDTO, ctx context.Context) error

	GetListPaginated(ctx context.Context,
		pagin *pagination.Paginator,
		filter *queryfilter.GetEventQueryFilter,
	) ([]EventUseCaseDTO, error)
}
