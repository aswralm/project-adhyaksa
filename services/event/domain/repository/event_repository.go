package repository

import (
	"context"
	"project-adhyaksa/pkg/pagination"
	"project-adhyaksa/services/event/domain/entity"
	queryfilter "project-adhyaksa/services/event/domain/query_filter"
)

type EventRepository interface {
	Create(event entity.Event, ctx context.Context) error

	//this method will using pagination and filter as response
	GetListPaginated(ctx context.Context,
		pagin *pagination.Paginator,
		filter *queryfilter.GetEventQueryFilter,
	) ([]entity.Event, error)
}
