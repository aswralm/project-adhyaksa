package model

import (
	"project-adhyaksa/services/event/domain/entity"
	"time"
)

type Event struct {
	ID          string     `dbq:"id" gorm:"primaryKey;column:id"`
	AdminID     string     `dbq:"admin_id" gorm:"column:admin_id"`
	BranchID    string     `dbq:"branch_id" gorm:"column:branch_id"`
	Name        string     `dbq:"name" gorm:"column:name"`
	StartTime   *time.Time `dbq:"start_time" gorm:"column:start_time"`
	EndTime     *time.Time `dbq:"end_time" gorm:"column:end_time"`
	Location    string     `dbq:"location" gorm:"column:location"`
	Description string     `dbq:"description" gorm:"column:description"`
	CreatedAt   time.Time  `dbq:"created_at" gorm:"column:created_at"`
	UpdatedAt   *time.Time `dbq:"updated_at" gorm:"column:updated_at"`
	DeletedAt   *time.Time `dbq:"deleted_at" gorm:"column:deleted_at"`

	Branch *Branch `gorm:"foreignkey:BranchID"`
}

func (Event) GetTableName() string {
	return "events"
}

func (m *Event) New(e entity.Event) *Event {

	m.ID = e.GetID()
	m.AdminID = e.GetAdminID()
	m.Name = e.GetName()
	m.StartTime = e.GetEndTime()
	m.EndTime = e.GetEndTime()
	m.Location = e.GetLocation()
	m.Description = e.GetDescription()
	m.BranchID = e.GetBranch().GetID()

	return m

}

func (m *Event) MapEventEntityList(events []Event) ([]*entity.Event, error) {
	entities := make([]*entity.Event, len(events))

	for i, model := range events {
		entity, err := MapEventEntity(&model)
		if err != nil {
			return nil, err
		}

		entities[i] = entity
	}

	return entities, nil
}

func MapEventEntity(modelEvent *Event) (*entity.Event, error) {

	branch, err := MapBranchEntity(modelEvent.Branch)
	if err != nil {
		return nil, err
	}

	entity, err := entity.NewEvent(entity.EventDTO{
		ID:          modelEvent.ID,
		Name:        modelEvent.Name,
		AdminID:     modelEvent.AdminID,
		StartTime:   modelEvent.StartTime,
		EndTime:     modelEvent.EndTime,
		Location:    modelEvent.Location,
		Description: modelEvent.Description,
		Organizer:   branch,
	})
	if err != nil {
		return nil, err
	}

	return entity, nil
}
