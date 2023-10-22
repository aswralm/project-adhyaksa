package service_test

import (
	"context"
	"errors"
	"project-adhyaksa/services/event/domain/service"
	"project-adhyaksa/services/event/internal/customerror"
	mocks "project-adhyaksa/services/event/internal/repository/mock"
	services "project-adhyaksa/services/event/internal/service"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	eventRepo = new(mocks.EventRepositoryMock)
)

func TestEventServiceCreate(t *testing.T) {
	var (
		startTime = time.Date(2023, 10, 10, 19, 0, 0, 0, time.UTC)
		endTime   = time.Date(2023, 10, 10, 22, 0, 0, 0, time.UTC)
	)
	testCases := []struct {
		name     string
		dto      service.EventServiceDTO
		expected any
		isError  bool
	}{
		{
			name: "positive case",
			dto: service.EventServiceDTO{
				Name:        "meeting tahunan",
				StartTime:   &startTime,
				EndTime:     &endTime,
				Location:    "jakarta",
				Description: "meningkatkan silatuirahmi",
				BranchID:    "cabang123",
			},
			expected: nil,
			isError:  false,
		},
		{
			name: "negative case: if name of event is empty",
			dto: service.EventServiceDTO{
				StartTime:   &startTime,
				EndTime:     &endTime,
				Location:    "jakarta",
				Description: "meningkatkan silatuirahmi",
				BranchID:    "cabang123",
			},
			expected: &customerror.Err{
				Code:   customerror.ERROR_INVALID_REQUEST,
				Errors: errors.New(customerror.ERROR_FIELD_ENTITY).Error(),
			},
			isError: true,
		},
		{
			name: "negative case: if branch_id of event is empty",
			dto: service.EventServiceDTO{
				Name:        "meeting tahunan",
				StartTime:   &startTime,
				EndTime:     &endTime,
				Location:    "jakarta",
				Description: "meningkatkan silatuirahmi",
			},
			expected: &customerror.Err{
				Code:   customerror.ERROR_INVALID_REQUEST,
				Errors: errors.New(customerror.ERROR_FIELD_ENTITY).Error(),
			},
			isError: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			eventRepo.Mock.On("Create", mock.Anything, mock.Anything).Return(nil)

			service := services.NewEventService(eventRepo)

			ctx := context.TODO()
			err := service.Create(testCase.dto, ctx)

			if testCase.isError {
				assert.Error(t, err)
				assert.Equal(t, testCase.expected, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
