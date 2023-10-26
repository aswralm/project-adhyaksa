package usecase_test

import (
	"context"
	usecases "project-adhyaksa/services/event/domain/usecase"
	valueobject "project-adhyaksa/services/event/domain/value_object"
	mocks "project-adhyaksa/services/event/internal/service/mock"
	"project-adhyaksa/services/event/internal/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	participantService = new(mocks.ParticipantServiceMock)
)

func TestParticipantServiceCreate(t *testing.T) {
	testCases := []struct {
		name     string
		dto      usecases.ParticipantUseCaseDTO
		expected any
		isError  bool
	}{
		{
			name: "positive case",
			dto: usecases.ParticipantUseCaseDTO{
				UserID:  "defb7588-55ed-45ea-a3ad-af1fcb37fb2a",
				EventID: "4cf68d55-aa88-48f1-b2ed-67873669a168",
				Status:  string(valueobject.StatusPresent),
			},
			expected: nil,
			isError:  false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			participantService.Mock.On("Create", mock.Anything, mock.Anything).Return(nil)

			usecase := usecase.NewParticipantUseCase(participantService)

			ctx := context.TODO()
			err := usecase.ConfirmEvent(testCase.dto, ctx)

			if testCase.isError {
				assert.Error(t, err)
				assert.Equal(t, testCase.expected, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
