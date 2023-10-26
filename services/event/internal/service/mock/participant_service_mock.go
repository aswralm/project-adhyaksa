package mocks

import (
	"context"
	"project-adhyaksa/services/event/domain/service"

	"github.com/stretchr/testify/mock"
)

type ParticipantServiceMock struct {
	Mock mock.Mock
}

func (r *ParticipantServiceMock) Create(participant service.ParticipantServiceDTO, ctx context.Context) error {
	args := r.Mock.Called(participant, ctx)
	return args.Error(0)
}
