package queries

import "project-adhyaksa/services/event/internal/repository/model"

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
