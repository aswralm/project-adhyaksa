package queries

import (
	"fmt"
	"project-adhyaksa/pkg/pagination"
	queryfilter "project-adhyaksa/services/event/domain/query_filter"
	"project-adhyaksa/services/event/internal/repository/model"
	"time"
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

func GetListEventFilter(
	event *model.Event,
	branch *model.Branch,
	pagin *pagination.Paginator,
	filter *queryfilter.GetEventQueryFilter,
) (query string, arg []any) {
	var (
		baseQuery = fmt.Sprintf("SELECT * FROM %s LEFT JOIN %s ON %s.branch_id = branchs.id", event.GetTableName(), branch.GetTableName(), event.GetTableName())
		arguments []any
	)
	if filter.NextEvent {
		baseQuery += " WHERE created_at >= ?"
		arguments = append(arguments, time.Now())
	}
	if filter.OrderBy != "" {
		baseQuery += " ORDER BY = ?"
		arguments = append(arguments, filter.OrderBy)
	}
	if filter.Limit > 0 {
		baseQuery += " LIMIT = ?"
		arguments = append(arguments, pagin.Limit)
	}
	if filter.Limit > 0 && filter.Page > 0 {
		baseQuery += "OFFSET = ?"
		arguments = append(arguments, pagin.Offset)
	}

	return baseQuery, arguments
}

func GetListEventCount(event model.Event) string {
	return fmt.Sprintf("SELECT COUNT(id) FROM %s", event.GetTableName())
}
