package mocks

import (
	"context"
	"project-adhyaksa/services/event/domain/service"

	"github.com/stretchr/testify/mock"
)

type EventServiceMock struct {
	Mock mock.Mock
}

func (r *EventServiceMock) Create(event service.EventServiceDTO, ctx context.Context) error {
	args := r.Mock.Called(event, ctx)
	return args.Error(0)
}
