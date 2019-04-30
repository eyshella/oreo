package calculation

import (
	"log"
	"oreo/models"
)

func FindIntersectionRayAndPolygon(start models.Point, direction models.Vector, polygon models.Polygon) models.Point {
	return models.Point{
		X: 0,
		Y: 0,
	}
}

func FindIntersectionOfRayAndSection(ray models.Vector, section models.Vector) *models.Point {
	if ray.Length() == 0 {
		log.Println("Error: Cannot find the intersection of ray and section, because ray direction length = 0")
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
	rayRatio := ((-section.From.X-ray.From.X)*(section.To.Y-section.From.Y) + (section.To.X-section.From.X)*(section.From.Y-ray.From.Y)) / determinant
	sectionRatio := ((ray.To.X-ray.From.X)*(section.From.Y-ray.From.Y) - (section.From.X-ray.From.X)*(ray.To.Y-ray.From.Y)) / determinant
	if sectionRatio < 0 || sectionRatio > 1 || rayRatio <= 0 {
		return nil
	}
	return &section.Multiply(sectionRatio).To
}
