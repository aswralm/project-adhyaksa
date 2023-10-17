package mocks

import (
	"context"
	"project-adhyaksa/pkg/pagination"
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

func (r *DocumentationRepositoryMock) GetListPaginated(pagin *pagination.Paginator, ctx context.Context) ([]*entity.Documentation, error) {
	args := r.Mock.Called(pagin, ctx)
	documentations, ok := args.Get(0).([]*entity.Documentation)
	if !ok {
		return nil, args.Error(1)
	}

	return documentations, args.Error(1)
}
