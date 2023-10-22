package usecase_test

import (
	"context"
	"project-adhyaksa/services/event/domain/usecase"
	mocks "project-adhyaksa/services/event/internal/service/mock"
	usecases "project-adhyaksa/services/event/internal/usecase"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	eventService = new(mocks.EventServiceMock)
)

func TestEventUseCaseCreate(t *testing.T) {
	var (
		startTime = time.Date(2023, 10, 10, 19, 0, 0, 0, time.UTC)
		endTime   = time.Date(2023, 10, 10, 22, 0, 0, 0, time.UTC)
	)
	testCases := []struct {
		name     string
		dto      usecase.EventUseCaseDTO
		expected any
		isError  bool
	}{
		{
			name: "positive case",
			dto: usecase.EventUseCaseDTO{
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
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			eventService.Mock.On("Create", mock.Anything, mock.Anything).Return(nil)

			useCase := usecases.NewEventUseCase(eventService)

			ctx := context.TODO()
			err := useCase.Create(testCase.dto, ctx)

			if testCase.isError {
				assert.Error(t, err)
				assert.Equal(t, testCase.expected, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
