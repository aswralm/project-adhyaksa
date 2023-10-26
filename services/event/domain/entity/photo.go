package entity

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

	result := &Photo{
		id:   photo.ID,
		url:  photo.URL,
		name: photo.Name,
	}

	if photo.Documentation != nil {
		result.documentation = photo.Documentation
	}
	return result, nil
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
