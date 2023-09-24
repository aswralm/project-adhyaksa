package entity_test

import (
	"project-adhyaksa/services/event/domain/entity"
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
				ID:      uuid.New().String(),
				UserID:  "user1",
				AdminID: "admin1",
				Status:  "active",
				Event:   &entity.Event{}, // assuming Event struct is defined in your entity package
			},
			expected: &expected{
				ID:      uuid.New().String(),
				UserID:  "user1",
				AdminID: "admin1",
				Status:  "active",
				Event:   &entity.Event{},
			},
			isError: false,
			err:     "",
		},
		{
			name: "negative case - missing ID",
			dto: entity.ParticipantDTO{
				ID:      "",
				UserID:  "user1",
				AdminID: "admin1",
				Status:  "active",
				Event:   &entity.Event{},
			},
			expected: nil,
			isError:  true,
			err:      "required field is missing",
		},
		{
			name: "negative case - missing Event",
			dto: entity.ParticipantDTO{
				ID:      uuid.New().String(),
				UserID:  "user1",
				AdminID: "admin1",
				Status:  "active",
				Event:   nil,
			},
			expected: nil,
			isError:  true,
			err:      "required field is missing",
		},
	}

	for _, test := range testcases {
		t.Run(test.name, func(t *testing.T) {

			participant, err := entity.NewParticipant(test.dto)

			if !test.isError {
				assert.NotNil(t, participant)
				assert.Equal(t, test.expected.UserID, participant.GetUserID())
				assert.Equal(t, test.expected.AdminID, participant.GetAdminID())
				assert.Equal(t, test.expected.Status, participant.GetStatus())
				assert.Equal(t, test.expected.Event, participant.GetEvent())
			} else {
				assert.Equal(t, test.err, err.Error())
				assert.Nil(t, participant)
			}

		})
	}
}
