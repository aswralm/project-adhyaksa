package request

import (
	queryfilter "project-adhyaksa/services/event/domain/query_filter"
)

type EventQueryPaginated struct {
	Limit     int    `form:"limit"`
	Page      int    `form:"page"`
	NextEvent bool   `form:"nextEvent"`
	Order     string `form:"order"`
	OrderBy   string `form:"orderBy"`
}

func (e *EventQueryPaginated) QueryParamMapping() queryfilter.GetEventQueryFilter {
	return queryfilter.GetEventQueryFilter{
		Limit:     e.Limit,
		Page:      e.Page,
		NextEvent: e.NextEvent,
		OrderBy:   e.OrderBy,
	}
}
