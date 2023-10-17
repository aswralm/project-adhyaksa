package queries

import (
	"project-adhyaksa/pkg/pagination"
	"project-adhyaksa/services/event/internal/repository/model"

	"gorm.io/gorm"
)

var RegisterDocumentationStatment = []string{"id", "admin_id", "branch_id", "name", "date", "participant", "location", "description", "created_at"}

func RegisterDocumentationArgument(documentationmodel *model.Documentation) []any {
	return []any{
		documentationmodel.ID,
		documentationmodel.AdminID,
		documentationmodel.BranchID,
		documentationmodel.Name,
		documentationmodel.Date,
		documentationmodel.Participant,
		documentationmodel.Location,
		documentationmodel.Description,
		documentationmodel.CreatedAt,
	}
}

func GetListDocumentationFilterGORM(pagin *pagination.Paginator) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		var (
			eventModel model.Event
		)
		db = db.Table(eventModel.GetTableName()).Preload("Branch").Preload("Photo")

		if pagin.Limit > 0 {
			db = db.Limit(pagin.Limit)
		}

		if pagin.Offset > 0 {
			db = db.Offset(pagin.Offset)
		}
		return db
	}
}

func GetListDocumentationCountGORM(db *gorm.DB) *gorm.DB {
	var (
		eventModel model.Event
	)
	db = db.Table(eventModel.GetTableName())

	return db
}
