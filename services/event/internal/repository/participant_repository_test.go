package repository_test

import (
	"context"
	"log"
	createid "project-adhyaksa/pkg/create-id"
	"project-adhyaksa/services/event/domain/entity"
	valueobject "project-adhyaksa/services/event/domain/value_object"
	"project-adhyaksa/services/event/internal/repository"
	"project-adhyaksa/services/event/internal/repository/model"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParticipantRepositoryCreate(t *testing.T) {
	config := SetupTest()

	//set refresh database
	var refresh = make(map[string]interface{})
	refresh[model.Participant{}.GetTableName()] = model.Participant{}
	refresh[model.Event{}.GetTableName()] = model.Event{}
	refresh[model.Branch{}.GetTableName()] = model.Branch{}

	refreshEventTable(config, refresh)
	//insert branch
	var data = make(map[string]interface{})
	data[model.Branch{}.GetTableName()] = model.Branch{
		ID:      "4cf68d55-aa88-48f1-b2ed-67873669a168",
		Name:    "branch test",
		Address: "jakarta",
	}

	startTime := time.Date(2023, 10, 10, 19, 0, 0, 0, time.UTC)
	endTime := time.Date(2023, 10, 10, 22, 0, 0, 0, time.UTC)
	data[model.Event{}.GetTableName()] = model.Event{
		ID:          "be53fd5c-3da6-4914-b870-35476e57cf0f",
		AdminID:     "defb7588-55ed-45ea-a3ad-af1fcb37fb2a",
		BranchID:    "4cf68d55-aa88-48f1-b2ed-67873669a168",
		Name:        "rakercap",
		StartTime:   &startTime,
		EndTime:     &endTime,
		Location:    "jakarta",
		Description: "lorem ipsum",
	}

	insertTable(config, data)

	//create entity
	eventEntity := entity.Event{}
	eventEntity.SetID("be53fd5c-3da6-4914-b870-35476e57cf0f")
	participantEntity, err := entity.NewParticipant(entity.ParticipantDTO{
		ID:     createid.CreateID(),
		UserID: createid.CreateID(),
		Status: string(valueobject.StatusPresent),
		Event:  &eventEntity,
	})
	if err != nil {
		log.Println(err)
	}

	testCases := []struct {
		name     string
		entity   *entity.Participant
		expected any
		isError  bool
	}{
		{
			name:     "positive case",
			entity:   participantEntity,
			expected: nil,
			isError:  false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			defer refreshEventTable(config, refresh)
			repository := repository.NewParticipantRepository(config)

			ctx := context.TODO()
			err := repository.Create(testCase.entity, ctx)

			if testCase.isError {
				assert.Error(t, err)
				assert.Equal(t, testCase.expected, err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
