package entity

type Photo struct {
	id   string
	url  string
	name string

	//relation
	documentation *Documentation
}
