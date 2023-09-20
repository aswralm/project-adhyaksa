package entity_test

import (
	"project-adhyaksa/services/event/domain/entity"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestPhotoEntity(t *testing.T) {
	type expected struct {
		ID   string
		URL  string
		Name string
	}
	type testCases struct {
		name     string
		dto      entity.PhotoDTO
		expected *expected
		isError  bool
		err      string
	}
	testcases := []testCases{
		{
			name: "positive case",
			dto: entity.PhotoDTO{
				ID:   uuid.New().String(),
				URL:  "https://example.com/photo1.jpg",
				Name: "photo1",
			},
			expected: &expected{
				ID:   uuid.New().String(),
				URL:  "https://example.com/photo1.jpg",
				Name: "photo1",
			},
			isError: false,
			err:     "",
		},
		{
			name: "negative case",
			dto: entity.PhotoDTO{
				ID:   "",
				URL:  "https://example.com/photo1.jpg",
				Name: "photo1",
			},
			expected: nil,
			isError:  true,
			err:      "ERROR_FIELD_ENTITY",
		},
	}

	for _, test := range testcases {
		t.Run(test.name, func(t *testing.T) {

			photo, err := entity.NewPhoto(test.dto)

			if !test.isError {
				assert.NotNil(t, photo.GetID())
				assert.Equal(t, test.expected.URL, photo.GetURL())
				assert.Equal(t, test.expected.Name, photo.GetName())
			} else {
				assert.Equal(t, test.err, err.Error())
				assert.Nil(t, photo)
			}

		})
	}
}
