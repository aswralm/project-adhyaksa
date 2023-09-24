package entity_test

import (
	"project-adhyaksa/services/event/domain/entity"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestDocumentationEntity(t *testing.T) {
	var (
		date, _ = time.Parse("2006-01-02 15:04:05", "2023-01-01 08:00:00")
	)

	type expected struct {
		ID          string
		AdminID     string
		Name        string
		Date        *time.Time
		Attendant   uint32
		Location    string
		Description string
		Participant uint32
	}
	type testCases struct {
		name     string
		dto      entity.DocumentationDTO
		expected *expected
		isError  bool
		err      string
	}
	testcases := []testCases{
		{
			name: "positive case",
			dto: entity.DocumentationDTO{
				ID:          uuid.New().String(),
				AdminID:     "Admin1",
				Name:        "Dokumentasi Pertemuan Bulanan",
				Date:        &date,
				Attendant:   100,
				Location:    "gedung serbaguna",
				Description: "membangun persatuan dan kesatuan",
				Participant: 80,
			},
			expected: &expected{
				ID:          uuid.New().String(),
				Name:        "Dokumentasi Pertemuan Bulanan",
				Date:        &date,
				Attendant:   100,
				Location:    "gedung serbaguna",
				Description: "membangun persatuan dan kesatuan",
				Participant: 80,
			},
			isError: false,
			err:     "",
		},
		{
			name: "negative case",
			dto: entity.DocumentationDTO{
				ID:          "",
				Name:        "Dokumentasi Pertemuan Bulanan",
				Date:        &date,
				Attendant:   100,
				Location:    "gedung serbaguna",
				Description: "membangun persatuan dan kesatuan",
				Participant: 80,
			},
			expected: nil,
			isError:  true,
			err:      "ERROR_FIELD_ENTITY",
		},
	}

	for _, test := range testcases {
		t.Run(test.name, func(t *testing.T) {

			documentation, err := entity.NewDocumentation(test.dto)

			if !test.isError {
				assert.NotNil(t, documentation)
			} else {
				assert.Nil(t, documentation)
				assert.Equal(t, test.err, err.Error())
			}
			if !test.isError {
				assert.NotNil(t, documentation)
			} else {
				assert.Nil(t, documentation)
				assert.Equal(t, test.err, err.Error())
			}

		})
	}
}
