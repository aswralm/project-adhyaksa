package request

import (
	"time"
)

type RegisterEventRequest struct {
	BranchID    string     `json:"branch_id" binding:"required"`
	Name        string     `json:"name" binding:"required"`
	StartTime   *time.Time `json:"start_time" binding:"required"`
	EndTime     *time.Time `json:"end_time" binding:"required"`
	Location    string     `json:"location" binding:"required"`
	Description string     `json:"description" binding:"required"`
}

func (reg *RegisterEventRequest) GetJsonFieldName(field string) string {
	return map[string]string{
		"BranchID":    "branch_id",
		"Name":        "name",
		"StartTime":   "start_time",
		"EndTime":     "end_time",
		"Location":    "location",
		"Description": "description",
	}[field]
}

// ErrMessages contains map of error messages
func (reg *RegisterEventRequest) ErrMessages() map[string]map[string]string {
	return map[string]map[string]string{
		"branch_id": {
			"required": "branch_id is required",
		},
		"name": {
			"required": "name is required",
		},
		"start_time": {
			"required": "start_time is required",
		},
		"end_time": {
			"required": "end_time is required",
		},
		"location": {
			"required": "location is required",
		},
		"description": {
			"required": "description is required",
		},
	}
}
