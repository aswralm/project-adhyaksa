package queries

import "project-adhyaksa/services/event/internal/repository/model"

var RegisterPhotoStatment = []string{"id", "documentation_id", "public_id", "url", "name", "created_at"}

func RegisterPhotoArgument(photomodel *model.Photo) []any {
	return []any{
		photomodel.ID,
		photomodel.DocumentationID,
		photomodel.PublicID,
		photomodel.URL,
		photomodel.Name,
		photomodel.CreatedAt,
	}
}
