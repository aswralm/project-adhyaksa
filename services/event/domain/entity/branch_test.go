package entity_test

import (
	"project-adhyaksa/services/event/domain/entity"
	"project-adhyaksa/services/event/internal/customerror"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestBranchEntity(t *testing.T) {
	type expected struct {
		Name    string
		ID      string
		Address string
	}
	type testCases struct {
		name     string
		dto      entity.BranchDTO
		expected *expected
		isError  bool
		err      string
	}
	testcases := []testCases{
		{
			name: "positive case",
			dto: entity.BranchDTO{
				ID:      uuid.New().String(),
				Name:    "Cabang Utama",
				Address: "Jl. Jendral Sudirman No.1, Makassar",
			},
			expected: &expected{
				Name:    "Cabang Utama",
				Address: "Jl. Jendral Sudirman No.1, Makassar",
			},
			isError: false,
			err:     "",
		},
		{
			name: "negative case",
			dto: entity.BranchDTO{
				ID:      "",
				Name:    "Cabang Utama",
				Address: "Jl. Jendral Sudirman No.1, Makassar",
			},
			expected: nil,
			isError:  true,
			err:      customerror.ERROR_INVALID_REQUEST,
		},
	}

	for _, test := range testcases {
		t.Run(test.name, func(t *testing.T) {

			branch, err := entity.NewBranch(test.dto)

			if !test.isError {
				assert.NotNil(t, branch.GetID())
				assert.Equal(t, test.expected.Name, branch.GetName())
				assert.Equal(t, test.expected.Address, branch.GetAddress())
			} else {
				assert.Equal(t, test.err, err.Error())
				assert.Nil(t, branch)
			}
		})
	}
}
