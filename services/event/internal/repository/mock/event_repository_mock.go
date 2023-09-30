package mocks

import (
	"context"
	"project-adhyaksa/services/event/domain/entity"

	"github.com/stretchr/testify/mock"
)

type EventRepositoryMock struct {
	Mock mock.Mock
}

func (r *EventRepositoryMock) Create(event entity.Event, ctx context.Context) error {
	args := r.Mock.Called(event, ctx)
	return args.Error(0)
}
