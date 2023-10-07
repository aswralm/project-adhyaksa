package queries

import (
	"project-adhyaksa/services/event/internal/repository/model"
)

var RegisterEventStatment = []string{"id", "admin_id", "branch_id", "name", "start_time", "end_time", "location", "description", "created_at", "updated_at", "deleted_at"}

func RegisterEventArgument(eventModel *model.Event) []any {
	return []any{
		eventModel.ID,
		eventModel.AdminID,
		eventModel.BranchID,
		eventModel.Name,
		eventModel.StartTime,
		eventModel.EndTime,
		eventModel.Location,
		eventModel.Description,
		eventModel.CreatedAt,
		eventModel.UpdatedAt,
		eventModel.DeletedAt,
	}
}
