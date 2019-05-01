package calculation

import (
	"math"
	"oreo/models"
)

func ReflectFromSection(ray models.Vector, section models.Vector) *models.Vector {
	intersection := FindIntersectionOfRayAndSection(ray, section)
	var result models.Vector
	if intersection == nil {
		return nil
	}
	ray = models.NewVector(models.NewPoint(ray.From.X-ray.To.X+intersection.X, ray.From.Y-ray.To.Y+intersection.Y), *intersection)
	if FindAngleBetweenVectors(ray, section) == math.Pi/2 {
		result = models.NewVector(ray.To, ray.From)
	} else if section.To.X == section.From.X {
		result = models.NewVector(*intersection, models.NewPoint(ray.From.X, ray.To.Y*2-ray.From.Y))
	} else if section.To.Y == section.From.Y {
		result = models.NewVector(*intersection, models.NewPoint(ray.To.X*2-ray.From.X, ray.From.Y))
	} else {
		k := (section.To.X - section.From.X) / (section.To.Y - section.From.Y)
		a := math.Pow(k, 2) + 1
		b := 2 * (k*(-k*ray.From.Y-ray.To.X+ray.From.X) - ray.To.Y)
		c := math.Pow(-k*ray.From.Y-ray.To.X+ray.From.X, 2) + ray.To.Y*ray.To.Y - math.Pow(ray.To.X-ray.From.X, 2) - math.Pow(ray.To.Y-ray.From.Y, 2)
		reflectedPointY := SolveQuadraticEquation(a, b, c)
		maxDistance := float64(0)
		for _, pointY := range reflectedPointY {
			point := models.NewPoint(k*(pointY-ray.From.Y)+ray.From.X, pointY)
			vector := models.NewVector(ray.From, point)
			distance := vector.Length()
			if distance > maxDistance {
				result = models.NewVector(ray.To, point)
				maxDistance = distance
			}
		}
	}
	return &result
}
