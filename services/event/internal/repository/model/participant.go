package model

import (
	"project-adhyaksa/services/event/domain/entity"
	"time"
)

type Participant struct {
	ID        string     `dbq:"id" gorm:"primaryKey;column:id"`
	UserID    string     `dbq:"user_id" gorm:"column:user_id"`
	EventID   string     `dbq:"event_id" gorm:"column:event_id"`
	Status    string     `dbq:"status" gorm:"column:status"`
	CreatedAt time.Time  `dbq:"created_at" gorm:"column:created_at"`
	UpdatedAt *time.Time `dbq:"updated_at" gorm:"column:updated_at"`
	DeletedAt *time.Time `dbq:"deleted_at" gorm:"column:deleted_at"`

	Branch *Event `gorm:"foreignkey:EventID"`
}

func (m Participant) GetTableName() string {
	return "participants"
}

func (m *Participant) New(e *entity.Participant) *Participant {

	m.ID = e.GetID()
	m.UserID = e.GetUserID()
	m.EventID = e.GetEvent().GetID()
	m.Status = string(e.GetStatus())

	return m

}
