package repository

type baseFilter struct {
	Limit  int
	Offset int
}
type ItemFilter struct {
	baseFilter
	PriceLeft  *int64
	PriceRight *int64
}

type OrderFilter struct {
	baseFilter
}
