package shop

type Order struct {
	Product   *Product
	SoldPrice float64 // keeping this, since price might change at different times
}
