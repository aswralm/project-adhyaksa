package usecase

import (
	"context"
	"mime/multipart"
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
	Location      string
	Description   string
	Participant   uint32
}

type DocumentatitonUseCase interface {
	Create(documentation DocumentationUseCaseDTO, file multipart.File, ctx context.Context) error
}
