package model

import (
	"project-adhyaksa/services/event/domain/entity"
	"time"
)

type Event struct {
	ID          string     `dbq:"id"`
	AdminID     string     `dbq:"admin_id"`
	BranchID    string     `dbq:"branch_id"`
	Name        string     `dbq:"name"`
	StartTime   *time.Time `dbq:"start_time"`
	EndTime     *time.Time `dbq:"end_time"`
	Location    string     `dbq:"location"`
	Description string     `dbq:"description"`
	CreatedAt   time.Time  `dbq:"created_at"`
	UpdatedAt   *time.Time `dbq:"updated_at"`
	DeletedAt   *time.Time `dbq:"deleted_at"`
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
