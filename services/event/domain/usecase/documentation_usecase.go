package usecase

import (
	"context"
	"time"
)

type DocumentationUseCaseDTO struct {
	ID            string
	BranchID      string
	AdminID       string
	Name          string
	PhotoID       string
	PhotoPublicID string
	PhotoURL      string
	PhotoName     string
	Date          *time.Time
	Attendant     uint32
	Location      string
	Description   string
	Participant   uint32
}

type DocumentatitonUseCase interface {
	Create(documentation DocumentationUseCaseDTO, ctx context.Context) error
}
