package repository_test

import (
	"context"
	"project-adhyaksa/services/event/domain/entity"
	"project-adhyaksa/services/event/internal/repository"
	"project-adhyaksa/services/event/internal/repository/model"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDocumentationRepositoryCreate(t *testing.T) {
	config := SetupTest()

	var refresh = make(map[string]interface{})
	refresh[model.Photo{}.GetTableName()] = model.Photo{}
	refresh[model.Documentation{}.GetTableName()] = model.Documentation{}
	refresh[model.Documentation{}.GetTableName()] = model.Documentation{}
	refresh[model.Branch{}.GetTableName()] = model.Branch{}

	refreshEventTable(config, refresh)

	//insert branch
	var data = make(map[string]interface{})
	data[model.Branch{}.GetTableName()] = model.Branch{
		ID:      "4cf68d55-aa88-48f1-b2ed-67873669a168",
		Name:    "branch test",
		Address: "jakarta",
	}

	insertTable(config, data)

	var (
		date                  = time.Date(2023, 10, 10, 19, 0, 0, 0, time.UTC)
		branchPositiveCase, _ = entity.NewBranch(entity.BranchDTO{
			ID: "4cf68d55-aa88-48f1-b2ed-67873669a168",
		})
	)

	documentationPositiveCase, _ := entity.NewDocumentation(entity.DocumentationDTO{
		ID:          "4cf68d55-aa88-48f1-b2ed-67873669a771",
		Name:        "documentation test",
		AdminID:     "test12",
		Date:        &date,
		Location:    "jakarta",
		Description: "event test",
		Branch:      branchPositiveCase,
	})

	photo, _ := entity.NewPhoto(entity.PhotoDTO{
		ID:            "4cf68d55-aa88-48f1-b2ed-67873669a122",
		PublicID:      "image_public_id",
		URL:           "www.image.gratis.test",
		Name:          "image-test.jpg",
		Documentation: documentationPositiveCase,
	})

	testCases := []struct {
		name     string
		entity1  *entity.Documentation
		entity2  *entity.Photo
		expected any
		isError  bool
	}{
		{
			name:     "positive case",
			entity1:  documentationPositiveCase,
			entity2:  photo,
			expected: nil,
			isError:  false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			defer refreshEventTable(config, refresh)
			repository := repository.NewDocumentationRepository(config)

			ctx := context.TODO()

			documentation := testCase.entity1
			photo := testCase.entity2

			err := repository.Create(*documentation, *photo, ctx)

			if testCase.isError {
				assert.Error(t, err)
				assert.Equal(t, testCase.expected, err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
