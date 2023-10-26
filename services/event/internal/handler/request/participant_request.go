package request

type ConfirmEventRequest struct {
	EventID string `json:"event_id" binding:"required"`
	Status  string `json:"status" binding:"required,oneof=absent present"`
}

func (r *ConfirmEventRequest) GetJsonFieldName(field string) string {
	return map[string]string{
		"EventID": "event_id",
		"Status":  "status",
	}[field]
}

// ErrMessages contains map of error messages
func (r *ConfirmEventRequest) ErrMessages() map[string]map[string]string {
	return map[string]map[string]string{
		"event_id": {
			"required": "event_id is required",
		},
		"status": {
			"required": "status is required",
			"oneof":    "status is absent or present only",
		},
	}
}
