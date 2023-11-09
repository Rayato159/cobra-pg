package src

import (
	"math"
)

const PI float64 = 3.1416

// Exceptions
const RadiusIsRequiredException string = "RadiusIsRequiredException"

func CalculateAreaCircle(r float64) float64 {
	return PI * math.Pow(r, 2)
}
