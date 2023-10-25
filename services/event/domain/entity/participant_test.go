package entity_test

import (
	"project-adhyaksa/services/event/domain/entity"
	"project-adhyaksa/services/event/internal/customerror"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestParticipantEntity(t *testing.T) {
	type expected struct {
		ID      string
		UserID  string
		AdminID string
		Status  string
		Event   *entity.Event
	}
	type testCases struct {
		name     string
		dto      entity.ParticipantDTO
		expected *expected
		isError  bool
		err      string
	}
	testcases := []testCases{
		{
			name: "positive case",
			dto: entity.ParticipantDTO{
				ID:     uuid.New().String(),
				UserID: "user1",
				Status: "present",
				Event:  &entity.Event{}, // assuming Event struct is defined in your entity package
			},
			expected: &expected{
				ID:     uuid.New().String(),
				UserID: "user1",
				Status: "present",
				Event:  &entity.Event{},
			},
			isError: false,
			err:     "",
		},
		{
			name: "negative case - missing status",
			dto: entity.ParticipantDTO{
				ID:     uuid.New().String(),
				UserID: "user1",
				Status: "",
				Event:  &entity.Event{},
			},
			expected: nil,
			isError:  true,
			err:      customerror.ERROR_INVALID_REQUEST,
		},
		{
			name: "negative case - missing Event",
			dto: entity.ParticipantDTO{
				ID:     uuid.New().String(),
				UserID: "user1",
				Status: "present",
				Event:  nil,
			},
			expected: nil,
			isError:  true,
			err:      customerror.ERROR_INVALID_REQUEST,
		},
	}

	for _, test := range testcases {
		t.Run(test.name, func(t *testing.T) {

			participant, err := entity.NewParticipant(test.dto)

			if !test.isError {
				assert.NotNil(t, participant)
				assert.Equal(t, test.expected.UserID, participant.GetUserID())
				assert.Equal(t, test.expected.Status, string(participant.GetStatus()))
				assert.Equal(t, test.expected.Event, participant.GetEvent())
			} else {
				assert.Equal(t, test.err, err.Error())
				assert.Nil(t, participant)
			}

		})
	}
}
