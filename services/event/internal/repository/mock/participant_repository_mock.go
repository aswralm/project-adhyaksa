package mocks

import (
	"context"
	"project-adhyaksa/services/event/domain/entity"

	"github.com/stretchr/testify/mock"
)

type ParticipantRepositoryMock struct {
	Mock mock.Mock
}

func (r *ParticipantRepositoryMock) Create(participant *entity.Participant, ctx context.Context) error {
	args := r.Mock.Called(participant, ctx)
	return args.Error(0)
}
