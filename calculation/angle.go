package calculation

import (
	"math"
	"oreo/models"
)

func AngleBetweenVectors(first models.Vector, second models.Vector) float64 {
	return math.Acos(first.ScalarMultiplyOnVector(second) / (first.Length() * second.Length()))
}
