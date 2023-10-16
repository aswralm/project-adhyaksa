package queryfilter

type GetEventQueryFilter struct {
	Limit     int
	Page      int
	NextEvent bool
	OrderBy   string
}
