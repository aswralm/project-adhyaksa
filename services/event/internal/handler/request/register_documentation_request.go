package request

import (
	"time"
)

type RegisterDocumentationRequest struct {
	BranchID    string     `form:"branch_id" binding:"required"`
	Name        string     `form:"name" binding:"required"`
	Date        *time.Time `form:"date" binding:"required"`
	Location    string     `form:"location" binding:"required"`
	Description string     `form:"description" binding:"required"`
	Participant uint32     `form:"participant" binding:"required"`
}

func (reg *RegisterDocumentationRequest) GetJsonFieldName(field string) string {
	return map[string]string{
		"BranchID":    "branch_id",
		"Name":        "name",
		"Date":        "date",
		"Location":    "location",
		"Description": "description",
		"Participant": "participant",
	}[field]
}

// ErrMessages contains map of error messages
func (reg *RegisterDocumentationRequest) ErrMessages() map[string]map[string]string {
	return map[string]map[string]string{
		"branch_id": {
			"required": "branch_id is required",
		},
		"name": {
			"required": "name is required",
		},
		"date": {
			"required": "date is required",
		},
		"location": {
			"required": "location is required",
		},
		"description": {
			"required": "description is required",
		},
		"participant": {
			"required": "participant is required",
		},
	}
}
