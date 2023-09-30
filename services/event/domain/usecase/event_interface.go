package usecase

import (
	"context"
	"time"
)

type EventUseCaseDTO struct {
	BranchID    string
	AdminID     string
	Name        string
	StartTime   *time.Time
	EndTime     *time.Time
	Location    string
	Description string
}

type EventUseCase interface {
	Create(event EventUseCaseDTO, ctx context.Context) error
}
