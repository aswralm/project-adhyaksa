package service_test

import (
	"context"
	"io"
	"project-adhyaksa/pkg/pagination"
	mock_upload "project-adhyaksa/pkg/upload/mock"
	"project-adhyaksa/services/event/domain/service"
	mocks "project-adhyaksa/services/event/internal/repository/mock"
	services "project-adhyaksa/services/event/internal/service"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	mockFile          = new(mocks.MockFile)
	documentationRepo = new(mocks.DocumentationRepositoryMock)
	cloudinaryUpload  = new(mock_upload.CloudinaryMock)
)

func TestDocumentationService_Create(t *testing.T) {
	var (
		date = time.Date(2023, 10, 10, 19, 0, 0, 0, time.UTC)
	)
	testCases := []struct {
		name     string
		dto      service.DocumentationServiceDTO
		expected any
		isError  bool
	}{
		{
			name: "positive case",
			dto: service.DocumentationServiceDTO{
				BranchID:    "branch123",
				Name:        "meeting tahunan",
				Date:        &date,
				Location:    "jakarta",
				Description: "meningkatkan silatuirahmi",
				Participant: 10,
			},
			expected: nil,
			isError:  false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			documentationRepo.Mock.On("Create", mock.Anything, mock.Anything, mock.Anything).Return(nil)

			cloudinaryUpload.Mock.On("UploadImage", mock.Anything, mock.Anything).Return("thisURL", "thisPublicID", nil)

			service := services.NewDocumentationService(documentationRepo, cloudinaryUpload)

			ctx := context.TODO()

			mockData := []byte("Hello, this is a mock file.")
			mockFile.On("Read", mock.Anything).Return(len(mockData), io.EOF).Once()

			err := service.Create(testCase.dto, mockFile, ctx)

			if testCase.isError {
				assert.Error(t, err)
				assert.Equal(t, testCase.expected, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestDocumentationService_GetListPaginated(t *testing.T) {
	testCases := []struct {
		name     string
		pagin    *pagination.Paginator
		expected []*service.DocumentationServiceDTO
		isError  bool
	}{
		{
			name: "positive case",
			pagin: &pagination.Paginator{
				Limit:  10,
				Offset: 0,
			},
			expected: []*service.DocumentationServiceDTO{},
			isError:  false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			documentationRepo.Mock.On("GetListPaginated", testCase.pagin, mock.Anything).Return(nil, nil)

			service := services.NewDocumentationService(documentationRepo, cloudinaryUpload)

			ctx := context.TODO()

			result, err := service.GetListPaginated(testCase.pagin, ctx)

			if testCase.isError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, testCase.expected, result)
			}
		})
	}
}
func TestDocumentationService_GetByID(t *testing.T) {
	testCases := []struct {
		name     string
		id       string
		expected *service.DocumentationServiceDTO
		isError  bool
	}{
		{
			name:     "positive case",
			id:       "doc123",
			expected: nil,
			isError:  false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			documentationRepo.Mock.On("GetByID", testCase.id, mock.Anything).Return(testCase.expected, nil)

			service := services.NewDocumentationService(documentationRepo, cloudinaryUpload)

			ctx := context.TODO()

			result, err := service.GetByID(testCase.id, ctx)

			if testCase.isError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, testCase.expected, result)
			}
		})
	}
}
