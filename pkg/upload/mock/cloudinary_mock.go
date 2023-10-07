package mock_upload

import (
	"context"
	"mime/multipart"

	"github.com/stretchr/testify/mock"
)

type CloudinaryMock struct {
	Mock mock.Mock
}

func (r *CloudinaryMock) UploadImage(ctx context.Context, file multipart.File) (url, publicID string, err error) {
	args := r.Mock.Called(ctx, file)
	return args.String(0), args.String(1), args.Error(2)
}

func (r *CloudinaryMock) RemoveImage(ctx context.Context, publicID string) error {
	args := r.Mock.Called(ctx, publicID)
	return args.Error(0)
}
