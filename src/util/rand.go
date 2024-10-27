package util

import "math/rand/v2"

func RandomFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}
