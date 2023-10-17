package request

type DocumentationQueryPaginated struct {
	Limit int `form:"limit"`
	Page  int `form:"page"`
}
