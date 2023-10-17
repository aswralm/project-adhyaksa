package model

import (
	"project-adhyaksa/services/event/domain/entity"
	"time"
)

type Documentation struct {
	ID          string     `dbq:"id" gorm:"primaryKey;column:id"`
	AdminID     string     `dbq:"admin_id" gorm:"column:admin_id"`
	BranchID    string     `dbq:"branch_id" gorm:"column:branch_id"`
	Name        string     `dbq:"name" gorm:"column:name"`
	Date        *time.Time `dbq:"date" gorm:"column:date"`
	Location    string     `dbq:"location" gorm:"column:location"`
	Description string     `dbq:"description" gorm:"column:description"`
	Participant uint32     `dbq:"participant" gorm:"column:participant"`
	CreatedAt   time.Time  `dbq:"created_at" gorm:"column:created_at"`
	UpdatedAt   *time.Time `dbq:"updated_at" gorm:"column:updated_at"`
	DeletedAt   *time.Time `dbq:"deleted_at" gorm:"column:deleted_at"`

	Branch *Branch `gorm:"foreignkey:BranchID"`
	Photos *[]Photo
}

func (Documentation) GetTableName() string {
	return "documentations"
}

func (m *Documentation) New(documentation entity.Documentation) *Documentation {
	return &Documentation{
		ID:          documentation.GetID(),
		AdminID:     documentation.GetAdminID(),
		BranchID:    documentation.GetBranch().GetID(),
		Name:        documentation.GetName(),
		Date:        documentation.GetDate(),
		Location:    documentation.GetLocation(),
		Description: documentation.GetDescription(),
		Participant: documentation.GetParticipant(),
	}
}

func (m *Documentation) MapDocumentationEntityList(documentations []Documentation) ([]*entity.Documentation, error) {
	entities := make([]*entity.Documentation, len(documentations))

	for i, model := range documentations {
		entity, err := MapDocumentationEntity(&model)
		if err != nil {
			return nil, err
		}
		entities[i] = entity
	}

	return entities, nil
}

func MapDocumentationEntity(documentation *Documentation) (*entity.Documentation, error) {

	branch, err := MapBranchEntity(documentation.Branch)
	if err != nil {
		return nil, err
	}

	var photoEntities = make([]entity.Photo, len(*documentation.Photos))
	for i, photoModel := range *documentation.Photos {
		photo, err := MapPhotoEntity(&photoModel)
		if err != nil {
			return nil, err
		}
		photoEntities[i] = *photo
	}

	entity, err := entity.NewDocumentation(entity.DocumentationDTO{
		ID:          documentation.ID,
		Name:        documentation.Name,
		AdminID:     documentation.AdminID,
		Date:        documentation.Date,
		Location:    documentation.Location,
		Description: documentation.Description,
		Branch:      branch,
		Photos:      &photoEntities,
	})
	if err != nil {
		return nil, err
	}

	return entity, nil
}
