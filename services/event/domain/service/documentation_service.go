package service

import (
	"context"
	"mime/multipart"
	"time"
)

type DocumentationServiceDTO struct {
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

type DocumentatitonService interface {
	Create(documentation DocumentationServiceDTO, file *multipart.File, ctx context.Context) error
}