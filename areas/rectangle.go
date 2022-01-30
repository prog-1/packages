package areas

// Rectangle returns the area of the rectangle with the given sides.
func Rectangle(a, b float64) float64 {
	if a < 0 || b < 0 {
		return 0
	}
	return a * b
}
