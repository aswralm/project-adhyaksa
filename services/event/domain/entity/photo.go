package entity

import (
	"errors"

	"github.com/google/uuid"
)

type Photo struct {
	id   string
	url  string
	name string

	//relation
	documentation *Documentation
}
type PhotoDTO struct {
	ID   string
	URL  string
	Name string

	//relation
	Documentation *Documentation
}

// mapping for DTO to Entity
func NewPhoto(photo PhotoDTO) (*Photo, error) {
	if photo.ID == "" || photo.URL == "" || photo.Name == "" {
		return nil, errors.New("ERROR_FIELD_ENTITY")
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
func (p Photo) SetID() {
	p.id = uuid.New().String()
}

func (p *Photo) GetID() string {
	return p.id
}

func (p Photo) SetURL(url string) {
	p.url = url
}

func (p *Photo) GetURL() string {
	return p.url
}

func (p Photo) SetName(name string) {
	p.name = name
}

func (p *Photo) GetName() string {
	return p.name
}

func (p Photo) SetDocumentation(documentation *Documentation) {
	p.documentation = documentation
}

func (p *Photo) GetDocumentation() *Documentation {
	return p.documentation
}
