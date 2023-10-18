package model

import (
	"project-adhyaksa/services/event/domain/entity"
	"time"
)

type Branch struct {
	ID        string     `dbq:"name" gorm:"primaryKey;column:id"`
	Name      string     `dbq:"id" gorm:"column:name"`
	Address   string     `dbq:"address" gorm:"column:address"`
	CreatedAt time.Time  `dbq:"created_at" gorm:"column:created_at"`
	UpdatedAt *time.Time `dbq:"updated_at" gorm:"column:updated_at"`
	DeletedAt *time.Time `dbq:"deleted_at" gorm:"column:deleted_at"`

	Event *[]Event
}

func (Branch) GetTableName() string {
	return "branches"
}

func MapBranchEntity(modelBranch *Branch) (*entity.Branch, error) {
	return entity.NewBranch(entity.BranchDTO{
		ID:      modelBranch.ID,
		Name:    modelBranch.Name,
		Address: modelBranch.Address,
	})
}
