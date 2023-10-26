package entity_test

import (
	"project-adhyaksa/services/event/domain/entity"
	"project-adhyaksa/services/event/internal/customerror"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestEventEntity(t *testing.T) {

	var (
		startTime, _ = time.Parse("2006-01-02 15:04:05", "2023-01-01 08:00:00")
		endTime, _   = time.Parse("2006-01-02 15:04:05", "2023-01-01 17:00:00")
	)

	type expected struct {
		Name        string
		StartTime   *time.Time
		EndTime     *time.Time
		Location    string
		Description string
	}
	type testCases struct {
		name     string
		dto      entity.EventDTO
		expected *expected
		isError  bool
		err      any
	}
	testcases := []testCases{
		{
			name: "positive case",
			dto: entity.EventDTO{
				ID:          uuid.New().String(),
				Name:        "Pertemuan Bulanan",
				StartTime:   &startTime,
				EndTime:     &endTime,
				Location:    "gedung serbaguna",
				Description: "membangun persatuan dan kesatuan",
			},
			expected: &expected{
				Name:        "Pertemuan Bulanan",
				StartTime:   &startTime,
				EndTime:     &endTime,
				Location:    "gedung serbaguna",
				Description: "membangun persatuan dan kesatuan",
			},
			isError: false,
			err:     nil,
		},
		{
			name: "negative case",
			dto: entity.EventDTO{
				ID:          uuid.New().String(),
				Name:        "",
				StartTime:   &startTime,
				EndTime:     &endTime,
				Location:    "gedung serbaguna",
				Description: "membangun persatuan dan kesatuan",
			},
			expected: nil,
			isError:  true,
			err:      customerror.ERROR_INVALID_REQUEST,
		},
	}

	for _, test := range testcases {
		t.Run(test.name, func(t *testing.T) {

			event, err := entity.NewEvent(test.dto)

			if !test.isError {
				assert.NotNil(t, event.GetID())
				assert.Equal(t, test.expected.Name, event.GetName())
				assert.Equal(t, test.expected.StartTime, event.GetStartTime())
				assert.Equal(t, test.expected.EndTime, event.GetEndTime())
				assert.Equal(t, test.expected.Location, event.GetLocation())
				assert.Equal(t, test.expected.Description, event.GetDescription())
			} else {
				assert.Equal(t, test.err, err.Error())
				assert.Nil(t, event)
			}
		})
	}
}
