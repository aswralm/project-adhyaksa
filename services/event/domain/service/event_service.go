package service

import (
	"context"
	"project-adhyaksa/pkg/pagination"
	queryfilter "project-adhyaksa/services/event/domain/query_filter"
	"time"
)

type EventServiceDTO struct {
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
type EventService interface {
	Create(event EventServiceDTO, ctx context.Context) error

	GetListPaginated(
		pagin *pagination.Paginator,
		filter *queryfilter.GetEventQueryFilter,
	) ([]EventServiceDTO, error)

	GetByID(id string, ctx context.Context) (*EventServiceDTO, error)
}
