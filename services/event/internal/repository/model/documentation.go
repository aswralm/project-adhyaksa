package model

import (
	"project-adhyaksa/services/event/domain/entity"
	"time"
)

type Documentation struct {
	ID          string     `dbq:"id"`
	AdminID     string     `dbq:"admin_id"`
	BranchID    string     `dbq:"branch_id"`
	Name        string     `dbq:"name"`
	Date        *time.Time `dbq:"date"`
	Location    string     `dbq:"location"`
	Description string     `dbq:"description"`
	Participant uint32     `dbq:"participant"`
	CreatedAt   time.Time  `dbq:"created_at"`
	UpdatedAt   *time.Time `dbq:"updated_at"`
	DeletedAt   *time.Time `dbq:"deleted_at"`
}

func (Documentation) GetTableName() string {
	return "documentations"
}

func (m *Documentation) New(documentation entity.Documentation) *Documentation {
	return &Documentation{
		ID:          documentation.GetID(),
		AdminID:     documentation.GetAdminID(),
		BranchID:    documentation.GetBranch().GetID(),
		Name:        documentation.GetName(),
		Date:        documentation.GetDate(),
		Location:    documentation.GetLocation(),
		Description: documentation.GetDescription(),
		Participant: documentation.GetParticipant(),
	}
}
