package usecase_test

import (
	"context"
	"project-adhyaksa/pkg/pagination"
	queryfilter "project-adhyaksa/services/event/domain/query_filter"
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

func TestEventUseCase_Create(t *testing.T) {
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

func TestEventUseCase_GetListPaginated(t *testing.T) {
	testCases := []struct {
		name     string
		pagin    *pagination.Paginator
		filter   *queryfilter.GetEventQueryFilter
		expected []usecase.EventUseCaseDTO
		isError  bool
	}{
		{
			name: "positive case",
			pagin: &pagination.Paginator{
				Limit:  10,
				Offset: 0,
			},
			filter:   &queryfilter.GetEventQueryFilter{},
			expected: []usecase.EventUseCaseDTO{},
			isError:  false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			eventService.Mock.On("GetListPaginated", mock.Anything, mock.Anything).Return([]usecase.EventUseCaseDTO{}, nil)

			useCase := usecases.NewEventUseCase(eventService)

			result, err := useCase.GetListPaginated(testCase.pagin, testCase.filter)

			if testCase.isError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, testCase.expected, result)
			}
		})
	}
}
