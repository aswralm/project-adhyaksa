package entity

import (
	"errors"
	"project-adhyaksa/services/event/internal/customerror"
)

type Photo struct {
	id       string
	publicID string
	url      string
	name     string

	//relation
	documentation *Documentation
}
type PhotoDTO struct {
	ID       string
	PublicID string
	URL      string
	Name     string

	//relation
	Documentation *Documentation
}

// mapping for DTO to Entity
func NewPhoto(photo PhotoDTO) (*Photo, error) {
	if photo.Name == "" {
		return nil, &customerror.Err{
			Code:   customerror.ERROR_INVALID_REQUEST,
			Errors: errors.New(customerror.ERROR_FIELD_ENTITY).Error(),
		}
	}

	return &Photo{
		id:   photo.ID,
		url:  photo.URL,
		name: photo.Name,

		//relation
		documentation: photo.Documentation,
	}, nil
}

// getter & setter for entity
func (p *Photo) SetID(id string) {
	p.id = id
}

func (p *Photo) GetID() string {
	return p.id
}

func (p *Photo) SetURL(url string) {
	p.url = url
}

func (p *Photo) GetURL() string {
	return p.url
}

func (p *Photo) SetName(name string) {
	p.name = name
}

func (p *Photo) GetName() string {
	return p.name
}

func (p *Photo) SetDocumentation(documentation *Documentation) {
	p.documentation = documentation
}

func (p *Photo) GetDocumentation() *Documentation {
	return p.documentation
}

func (p *Photo) SetPublicID(publicID string) {
	p.publicID = publicID
}

func (p *Photo) GetPublicID() string {
	return p.publicID
}
