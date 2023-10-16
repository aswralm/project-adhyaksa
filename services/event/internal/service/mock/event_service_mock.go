package mocks

import (
	"context"
	"project-adhyaksa/pkg/pagination"
	"project-adhyaksa/services/event/domain/service"

	queryfilter "project-adhyaksa/services/event/domain/query_filter"

	"github.com/stretchr/testify/mock"
)

type EventServiceMock struct {
	Mock mock.Mock
}

func (r *EventServiceMock) Create(event service.EventServiceDTO, ctx context.Context) error {
	args := r.Mock.Called(event, ctx)
	return args.Error(0)
}

func (r *EventServiceMock) GetListPaginated(
	pagin *pagination.Paginator,
	filter *queryfilter.GetEventQueryFilter,
) ([]service.EventServiceDTO, error) {
	args := r.Mock.Called(pagin, filter)
	events, ok := args.Get(0).([]service.EventServiceDTO)
	if !ok {
		return nil, args.Error(1)
	}

	return events, args.Error(1)
}

func (r *EventServiceMock) GetByID(id string, ctx context.Context) (*service.EventServiceDTO, error) {
	args := r.Mock.Called(id, ctx)
	event, ok := args.Get(0).(*service.EventServiceDTO)
	if !ok {
		return nil, args.Error(1)
	}

	return event, args.Error(1)
}
