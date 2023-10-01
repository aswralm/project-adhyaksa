package model

import (
	"project-adhyaksa/services/event/domain/entity"
	"time"
)

type Photo struct {
	ID              string     `dbq:"id"`
	DocumentationID string     `dbq:"documentation_id"`
	PublicID        string     `dbq:"public_id"`
	URL             string     `dbq:"url"`
	Name            string     `dbq:"name"`
	CreatedAt       time.Time  `dbq:"created_at"`
	UpdatedAt       *time.Time `dbq:"updated_at"`
	DeletedAt       *time.Time `dbq:"deleted_at"`
}

func (Photo) GetTableName() string {
	return "photos"
}

func (m *Photo) New(photo entity.Photo) *Photo {
	return &Photo{
		ID:              photo.GetID(),
		DocumentationID: photo.GetDocumentation().GetID(),
		PublicID:        photo.GetPublicID(),
		URL:             photo.GetURL(),
		Name:            photo.GetName(),
	}
}
