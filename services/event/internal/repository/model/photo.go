package model

import (
	"project-adhyaksa/services/event/domain/entity"
	"time"
)

type Photo struct {
	ID              string     `dbq:"id" gorm:"primaryKey;column:id"`
	DocumentationID string     `dbq:"documentation_id" gorm:"column:documentation_id"`
	PublicID        string     `dbq:"public_id" gorm:"column:public_id"`
	URL             string     `dbq:"url" gorm:"column:url"`
	Name            string     `dbq:"name" gorm:"column:name"`
	CreatedAt       time.Time  `dbq:"created_at" gorm:"column:created_at"`
	UpdatedAt       *time.Time `dbq:"updated_at" gorm:"column:updated_at"`
	DeletedAt       *time.Time `dbq:"deleted_at" gorm:"column:deleted_at"`

	Documentation *Documentation `gorm:"foreignkey:DocumentationID"`
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

func MapPhotoEntity(modelPhoto *Photo) (*entity.Photo, error) {
	return entity.NewPhoto(entity.PhotoDTO{
		ID:       modelPhoto.ID,
		PublicID: modelPhoto.PublicID,
		URL:      modelPhoto.URL,
		Name:     modelPhoto.Name,
	})
}
