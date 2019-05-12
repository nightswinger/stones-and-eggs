package intarray

import "math"

func Min(a []int) int {
	r := float64(a[0])
	for i := 0; i < len(a); i++ {
		r = math.Min(r, float64(a[i]))
	}

	return int(r)
}

func Max(a []int) int {
	r := float64(a[0])
	for i := 0; i < len(a); i++ {
		r = math.Max(r, float64(a[i]))
	}

	return int(r)
}
