package entity

import (
	"errors"

	"github.com/google/uuid"
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
		return nil, errors.New("ERROR_FIELD_ENTITY")
	}

	return &Branch{
		id:      branch.ID,
		name:    branch.Name,
		address: branch.Address,
	}, nil
}

// getter & setter for entity
func (b Branch) SetID() {
	b.id = uuid.New().String()
}

func (b *Branch) GetID() string {
	return b.id
}

func (b Branch) SetName(name string) {
	b.name = name
}

func (b *Branch) GetName() string {
	return b.name
}

func (b Branch) SetAddress(address string) {
	b.address = address
}

func (b *Branch) GetAddress() string {
	return b.address
}
