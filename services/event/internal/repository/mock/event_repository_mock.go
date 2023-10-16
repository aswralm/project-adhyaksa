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

func (r *EventRepositoryMock) GetListPaginated(
	pagin *pagination.Paginator,
	filter *queryfilter.GetEventQueryFilter,
) ([]*entity.Event, error) {
	args := r.Mock.Called(pagin, filter)
	events, ok := args.Get(0).([]*entity.Event)
	if !ok {
		return nil, args.Error(1)
	}

	return events, args.Error(1)
}

func (r *EventRepositoryMock) GetByID(id string, ctx context.Context) (*entity.Event, error) {
	args := r.Mock.Called(id, ctx)
	event, ok := args.Get(0).(*entity.Event)
	if !ok {
		return nil, args.Error(1)
	}

	return event, args.Error(1)
}
