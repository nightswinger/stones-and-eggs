package mathutil

import "math"

// Round rounds a float
func Round(num, precision float64) float64 {
	shift := math.Pow(10, precision)
	return roundInt(num * shift)
}

func roundInt(num float64) float64 {
	t := math.Trunc(num)
	if math.Abs(num-t) >= 0.5 {
		return t + math.Copysign(1, num)
	}

	return t
}
