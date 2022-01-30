package areas

import "math"

// Circle returns the area of the circle with the given radius.
func Circle(r float64) float64 {
	if r < 0 {
		return 0
	}
	return math.Pi * r * r
}
