package calculation

import (
	"log"
	"oreo/config"
	"oreo/models"
)

func FindIntersectionRayAndPolygon(ray models.Vector, polygon models.Polygon) (*models.Point, int, int) {
	log.Printf("FindIntersectionRayAndPolygon started. Params: ray: %s, polygon: %s", ray.ToString(), polygon.ToString())
	for k := 0; k < len(polygon.Vertices)-1; k++ {
		intersection := FindIntersectionOfRayAndSection(ray, models.NewVector(polygon.Vertices[k], polygon.Vertices[k+1]))
		if intersection != nil && !ray.From.Equal(*intersection) {
			log.Printf("FindIntersectionRayAndPolygon finished. Return: %s, %d, %d", intersection.ToString(), k, k+1)
			return intersection, k, k + 1
		}
	}
	intersection := FindIntersectionOfRayAndSection(ray, models.NewVector(polygon.Vertices[len(polygon.Vertices)-1], polygon.Vertices[0]))
	if intersection != nil && !ray.From.Equal(*intersection) {
		log.Printf("FindIntersectionRayAndPolygon finished. Return: %s, %d, %d", intersection.ToString(), len(polygon.Vertices)-1, 0)
		return intersection, len(polygon.Vertices) - 1, 0
	}
	log.Printf("FindIntersectionRayAndPolygon finished. Return: nill, -1, -1")
	return nil, -1, -1
}

func FindIntersectionOfRayAndSection(ray models.Vector, section models.Vector) *models.Point {
	log.Printf("FindIntersectionOfRayAndSection started. Params: ray: %s, section: %s", ray.ToString(), section.ToString())
	if ray.Length() == 0 {
		log.Println("FindIntersectionOfRayAndSection error. Cannot find the intersection of ray and section, because ray direction length = 0")
		return nil
	}
	if section.Length() == 0 {
		if ray.To.X-ray.From.X != 0 {
			c := (section.From.X - ray.From.X) / (ray.To.X - ray.From.X)
			if c*(ray.To.Y-ray.From.Y) == (section.From.Y - ray.From.Y) {
				return &section.From
			}
		}
		if ray.To.Y-ray.From.Y != 0 {
			c := (section.From.Y - ray.From.Y) / (ray.To.Y - ray.From.Y)
			if c*(ray.To.X-ray.From.X) == (section.From.X - ray.From.X) {
				return &section.From
			}
		}
		return nil
	}
	//Solving system of linear equations by Kramer method
	determinant := -(ray.To.X-ray.From.X)*(section.To.Y-section.From.Y) + (section.To.X-section.From.X)*(ray.To.Y-ray.From.Y)
	if determinant == 0 {
		return nil
	}
	rayRatio := (-(section.From.X-ray.From.X)*(section.To.Y-section.From.Y) + (section.To.X-section.From.X)*(section.From.Y-ray.From.Y)) / determinant
	sectionRatio := ((ray.To.X-ray.From.X)*(section.From.Y-ray.From.Y) - (section.From.X-ray.From.X)*(ray.To.Y-ray.From.Y)) / determinant
	log.Printf("FindIntersectionOfRayAndSection. rayRatio: %f, sectionRatio: %f", rayRatio, sectionRatio)
	if sectionRatio+config.Config.Accuracy < 0 || sectionRatio-config.Config.Accuracy > 1 || rayRatio+config.Config.Accuracy <= 0 {
		log.Printf("FindIntersectionOfRayAndSection finished. Return: nill")
		return nil
	}
	result := section.MultiplyOnScalar(sectionRatio).To
	log.Printf("FindIntersectionOfRayAndSection finished. Return: %s", result.ToString())
	return &result
}
