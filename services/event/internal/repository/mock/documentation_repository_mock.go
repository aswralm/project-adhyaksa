package mocks

import (
	"context"
	"project-adhyaksa/services/event/domain/entity"

	"github.com/stretchr/testify/mock"
)

type DocumentationRepositoryMock struct {
	Mock mock.Mock
}

func (r *DocumentationRepositoryMock) Create(documentation entity.Documentation, photo entity.Photo, ctx context.Context) error {
	args := r.Mock.Called(documentation, photo, ctx)
	return args.Error(0)
}
