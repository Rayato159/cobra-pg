package src

import (
	"math"
)

// Exceptions
const SquareDimensionIsRequiredException string = "SquareDimensionIsRequiredException"

func CalculateAreaSquare(d float64) float64 {
	return math.Pow(d, 2)
}
