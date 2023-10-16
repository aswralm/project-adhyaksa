package entity

import (
	"errors"
	"project-adhyaksa/services/event/internal/customerror"
)

type Branch struct {
	name    string
	id      string
	address string
}
type BranchDTO struct {
	Name    string
	ID      string
	Address string
}

// mapping for DTO to Entity
func NewBranch(branch BranchDTO) (*Branch, error) {
	if branch.ID == "" {
		return nil, &customerror.Err{
			Code:   customerror.ERROR_INVALID_REQUEST,
			Errors: errors.New(customerror.ERROR_FIELD_ENTITY).Error(),
		}
	}

	return &Branch{
		id:      branch.ID,
		name:    branch.Name,
		address: branch.Address,
	}, nil
}

// getter & setter for entity
func (b *Branch) SetID(id string) {
	b.id = id
}

func (b *Branch) GetID() string {
	return b.id
}

func (b *Branch) SetName(name string) {
	b.name = name
}

func (b *Branch) GetName() string {
	return b.name
}

func (b *Branch) SetAddress(address string) {
	b.address = address
}

func (b *Branch) GetAddress() string {
	return b.address
}
