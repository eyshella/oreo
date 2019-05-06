package calculation

import (
	"log"
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
	n := models.NewVector(*intersection, models.NewPoint(section.From.Y-section.To.Y+intersection.X, section.To.X-section.From.X+intersection.Y))
	n = n.MultiplyOnScalar(1 / n.Length())
	result = ray.Subtract(n.MultiplyOnScalar(2 * ray.ScalarMultiplyOnVector(n)))
	result = models.NewVector(*intersection, models.NewPoint(result.To.X-result.From.X+intersection.X, result.To.Y-result.From.Y+intersection.Y))
	log.Printf("ReflectFromSection finished. Return: %s", result.ToString())
	return &result
}

func ReflectFromPolygonVertice(ray models.Vector, polygon models.Polygon, k int) *models.Vector {
	log.Printf("ReflectFromPolygonVertice started. Params: ray: %s, polygon: %s, verticle number: %d", ray.ToString(), polygon.ToString(), k+1)
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
