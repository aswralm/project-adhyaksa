package service_test

import (
	"context"
	"errors"
	createid "project-adhyaksa/pkg/create-id"
	"project-adhyaksa/services/event/domain/service"
	valueobject "project-adhyaksa/services/event/domain/value_object"
	"project-adhyaksa/services/event/internal/customerror"
	mocks "project-adhyaksa/services/event/internal/repository/mock"
	services "project-adhyaksa/services/event/internal/service"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	participantRepo = new(mocks.ParticipantRepositoryMock)
)

func TestParticipantServiceCreate(t *testing.T) {
	testCases := []struct {
		name     string
		dto      service.ParticipantServiceDTO
		expected any
		isError  bool
	}{
		{
			name: "positive case",
			dto: service.ParticipantServiceDTO{
				ID:      createid.CreateID(),
				UserID:  "defb7588-55ed-45ea-a3ad-af1fcb37fb2a",
				EventID: "4cf68d55-aa88-48f1-b2ed-67873669a168",
				Status:  string(valueobject.StatusPresent),
			},
			expected: nil,
			isError:  false,
		},
		{
			name: "negative case: if status is empty",
			dto: service.ParticipantServiceDTO{
				ID:      createid.CreateID(),
				UserID:  "defb7588-55ed-45ea-a3ad-af1fcb37fb2a",
				EventID: "4cf68d55-aa88-48f1-b2ed-67873669a168",
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
			participantRepo.Mock.On("Create", mock.Anything, mock.Anything).Return(nil)

			service := services.NewParticipantService(participantRepo)

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
