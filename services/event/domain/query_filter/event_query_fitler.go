package queryfilter

type GetEventQueryFilter struct {
	Limit     int
	Page      int
	NextEvent bool
	Order     string
	OrderBy   string
}
