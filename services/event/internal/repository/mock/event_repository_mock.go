package mocks

import (
	"context"
	"project-adhyaksa/pkg/pagination"
	"project-adhyaksa/services/event/domain/entity"
	queryfilter "project-adhyaksa/services/event/domain/query_filter"

	"github.com/stretchr/testify/mock"
)

type EventRepositoryMock struct {
	Mock mock.Mock
}

func (r *EventRepositoryMock) Create(event entity.Event, ctx context.Context) error {
	args := r.Mock.Called(event, ctx)
	return args.Error(0)
}

func (r *EventRepositoryMock) GetListPaginated(ctx context.Context,
	pagin *pagination.Paginator,
	filter *queryfilter.GetEventQueryFilter,
) ([]*entity.Event, error) {
	args := r.Mock.Called(ctx, pagin, filter)
	events, ok := args.Get(0).([]*entity.Event)
	if !ok {
		return nil, args.Error(1)
	}

	return events, args.Error(1)
}
