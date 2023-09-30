package service

import (
	"context"
	"time"
)

type EventServiceDTO struct {
	BranchID    string
	AdminID     string
	Name        string
	StartTime   *time.Time
	EndTime     *time.Time
	Location    string
	Description string
}
type EventService interface {
	Create(event EventServiceDTO, ctx context.Context) error
}
