package mocks

import (
	"context"
	"mime/multipart"
	"project-adhyaksa/services/event/domain/service"

	"github.com/stretchr/testify/mock"
)

type DocumentationServiceMock struct {
	Mock mock.Mock
}

func (r *DocumentationServiceMock) Create(documentation service.DocumentationServiceDTO, file multipart.File, ctx context.Context) error {
	args := r.Mock.Called(documentation, file, ctx)
	return args.Error(0)
}
