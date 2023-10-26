package repository_test

import (
	"context"
	createid "project-adhyaksa/pkg/create-id"
	"project-adhyaksa/pkg/pagination"
	"project-adhyaksa/services/event/domain/entity"
	queryfilter "project-adhyaksa/services/event/domain/query_filter"
	"project-adhyaksa/services/event/internal/repository"
	"project-adhyaksa/services/event/internal/repository/model"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestEventRepositoryCreate(t *testing.T) {
	config := SetupTest()
	//insert branch
	var data = make(map[string]interface{})
	data[model.Branch{}.GetTableName()] = model.Branch{
		ID:      "4cf68d55-aa88-48f1-b2ed-67873669a168",
		Name:    "branch test",
		Address: "jakarta",
	}

	insertTable(config, data)

	var refresh = make(map[string]interface{})
	refresh[model.Event{}.GetTableName()] = model.Event{}
	refresh[model.Photo{}.GetTableName()] = model.Photo{}
	refresh[model.Documentation{}.GetTableName()] = model.Documentation{}
	refresh[model.Branch{}.GetTableName()] = model.Branch{}

	var (
		startTime             = time.Date(2023, 10, 10, 19, 0, 0, 0, time.UTC)
		endTime               = time.Date(2023, 10, 10, 22, 0, 0, 0, time.UTC)
		branchPositiveCase, _ = entity.NewBranch(entity.BranchDTO{
			ID: "4cf68d55-aa88-48f1-b2ed-67873669a168",
		})
		eventPositiveCase, _ = entity.NewEvent(entity.EventDTO{
			ID:          createid.CreateID(),
			Name:        "test",
			AdminID:     "test12",
			StartTime:   &startTime,
			EndTime:     &endTime,
			Location:    "jakarta",
			Description: "event test",
			Organizer:   branchPositiveCase,
		})

		branchNegativeCase, _ = entity.NewBranch(entity.BranchDTO{
			ID: "4cf68d55-aa88-48f1-b2ed-67873669a168",
		})
		eventNegativeCase, _ = entity.NewEvent(entity.EventDTO{
			ID:          createid.CreateID(),
			Name:        "test",
			AdminID:     "test12",
			StartTime:   &startTime,
			EndTime:     &endTime,
			Location:    "jakarta",
			Description: "event test",
			Organizer:   branchNegativeCase,
		})
	)
	testCases := []struct {
		name     string
		entity   entity.Event
		expected any
		isError  bool
	}{
		{
			name:     "positive case",
			entity:   *eventPositiveCase,
			expected: nil,
			isError:  false,
		},
		{
			name:     "negative case",
			entity:   *eventNegativeCase,
			expected: string("Error 1452 (23000): Cannot add or update a child row: a foreign key constraint fails (`event_sevice_test`.`events`, CONSTRAINT `events_ibfk_1` FOREIGN KEY (`branch_id`) REFERENCES `branches` (`id`))"),
			isError:  true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			defer refreshEventTable(config, refresh)
			repository := repository.NewEventRepository(config)

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

func TestEventRepositoryGetListPaginate(t *testing.T) {
	config := SetupTest()
	//insert branch
	var data = make(map[string]interface{})
	data[model.Branch{}.GetTableName()] = model.Branch{
		ID:      "4cf68d55-aa88-48f1-b2ed-67873669a168",
		Name:    "branch 12355",
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

	var refresh = make(map[string]interface{})
	refresh[model.Event{}.GetTableName()] = model.Event{}
	refresh[model.Photo{}.GetTableName()] = model.Photo{}
	refresh[model.Documentation{}.GetTableName()] = model.Documentation{}
	refresh[model.Branch{}.GetTableName()] = model.Branch{}

	testCases := []struct {
		name     string
		pagin    pagination.Paginator
		filter   queryfilter.GetEventQueryFilter
		expected any
		isError  bool
	}{
		{
			name:     "positive case",
			expected: nil,
			isError:  false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			defer refreshEventTable(config, refresh)
			repository := repository.NewEventRepository(config)

			_, err := repository.GetListPaginated(&testCase.pagin, &testCase.filter)

			if testCase.isError {
				assert.Error(t, err)
				assert.Equal(t, testCase.expected, err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestEventRepositoryGetByID(t *testing.T) {
	config := SetupTest()
	//insert branch
	var data = make(map[string]interface{})
	data[model.Branch{}.GetTableName()] = model.Branch{
		ID:      "4cf68d55-aa88-48f1-b2ed-67873669a168",
		Name:    "branch 12355",
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

	var refresh = make(map[string]interface{})
	refresh[model.Event{}.GetTableName()] = model.Event{}
	refresh[model.Photo{}.GetTableName()] = model.Photo{}
	refresh[model.Documentation{}.GetTableName()] = model.Documentation{}
	refresh[model.Branch{}.GetTableName()] = model.Branch{}

	testCases := []struct {
		name     string
		pagin    pagination.Paginator
		filter   queryfilter.GetEventQueryFilter
		expected any
		isError  bool
	}{
		{
			name:     "positive case",
			expected: nil,
			isError:  false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			defer refreshEventTable(config, refresh)
			repository := repository.NewEventRepository(config)
			ctx := context.TODO()
			_, err := repository.GetByID("be53fd5c-3da6-4914-b870-35476e57cf0f", ctx)

			if testCase.isError {
				assert.Error(t, err)
				assert.Equal(t, testCase.expected, err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
