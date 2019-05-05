package calculation

import (
	"log"
	"math"
	"oreo/models"
)

func ReflectFromSection(ray models.Vector, section models.Vector) *models.Vector {
	log.Printf("ReflectFromSection started. Params: ray: %s, section: %s", ray.ToString(), section.ToString())
	intersection := FindIntersectionOfRayAndSection(ray, section)
	var result models.Vector
	if intersection == nil {
		log.Printf("ReflectFromSection finished. Return: nill")
		return nil
	}
	ray = models.NewVector(models.NewPoint(ray.From.X-ray.To.X+intersection.X, ray.From.Y-ray.To.Y+intersection.Y), *intersection)
	if section.To.X == section.From.X {
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
			if distance >= maxDistance {
				result = models.NewVector(ray.To, point)
				maxDistance = distance
			}
		}
	}
	log.Printf("ReflectFromSection finished. Return: %s", result.ToString())
	return &result
}

func ReflectFromPolygonVertice(ray models.Vector, polygon models.Polygon, k int) *models.Vector {
	log.Printf("ReflectFromPolygonVertice started. Params: ray: %s, polygon: %s, k: %d", ray.ToString(), polygon.ToString(), k)
	var previousSection models.Vector
	var nextSection models.Vector
	vertice := polygon.Vertices[k]
	if k == 0 {
		previousSection = models.NewVector(vertice, polygon.Vertices[len(polygon.Vertices)-1])
		nextSection = models.NewVector(vertice, polygon.Vertices[1])
	} else if k == len(polygon.Vertices)-1 {
		previousSection = models.NewVector(vertice, polygon.Vertices[k-1])
		nextSection = models.NewVector(vertice, polygon.Vertices[0])
	} else {
		previousSection = models.NewVector(vertice, polygon.Vertices[k-1])
		nextSection = models.NewVector(vertice, polygon.Vertices[k+1])
	}
	result := ReflectFromSection(ray, models.NewVector(models.NewPoint(previousSection.To.X-nextSection.To.X+vertice.X, previousSection.To.Y-nextSection.To.Y+vertice.Y), models.NewPoint(nextSection.To.X-previousSection.To.X+vertice.X, nextSection.To.Y-previousSection.To.Y+vertice.Y)))
	if result != nil {
		log.Printf("ReflectFromPolygonVertice finished. Return: %s", result.ToString())
	} else {
		log.Printf("ReflectFromPolygonVertice finished. Return: nil")
	}
	return result
}
