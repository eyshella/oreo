package calculation

import (
	"log"
	"math/rand"
	"oreo/models"
	"time"
)

func Step(start models.Point, polygon models.Polygon, way int, alpha float64) (models.Point, int) {
	verticleNumber := way
	direction := models.NewVector(start, polygon.Vertices[way])
	distance := direction.Length()
	path := distance * alpha
	intersections := 0
	isReflectionFromVerticle := true
	sectionFrom := -1
	sectionTo := -1
	for path > 0 {
		if path <= distance {
			log.Println("path <= distance")
			direction = *direction.MultiplyOnScalar(path / distance)
			path = 0
		} else if path > distance {
			log.Println("path > distance")
			intersections++
			path -= distance
			var newDirection *models.Vector
			if isReflectionFromVerticle {
				newDirection = ReflectFromPolygonVertice(direction, polygon, verticleNumber)
				log.Println("Reflection from verticle")
				log.Printf("Verticle: %s", polygon.Vertices[verticleNumber].ToString())
			} else {
				log.Println("Reflection from section")
				newDirection = ReflectFromSection(direction, models.NewVector(polygon.Vertices[sectionFrom], polygon.Vertices[sectionTo]))
			}
			from := direction.To
			intersection, sectionFrom, sectionTo := FindIntersectionRayAndPolygon(*newDirection, polygon)
			direction = models.NewVector(from, *intersection)
			distance = direction.Length()
			isReflectionFromVerticle = false
			if polygon.Vertices[sectionFrom].Equal(direction.To) {
				isReflectionFromVerticle = true
				verticleNumber = sectionFrom
			}
			if polygon.Vertices[sectionTo].Equal(direction.To) {
				isReflectionFromVerticle = true
				verticleNumber = sectionTo
			}
		}
	}
	return direction.To, intersections
}

func Walk(start models.Point, polygon models.Polygon, steps int) ([]models.Point, int) {
	rand.Seed(time.Now().UnixNano())
	point := start
	points := make([]models.Point, 0)
	points = append(points, point)
	intersections := 0
	for i := 0; i < steps; i++ {
		alpha := 1.5
		way := rand.Intn(len(polygon.Vertices))
		log.Println("New step")
		point, stepIntersections := Step(point, polygon, way, alpha)
		log.Printf("Point: %s", point.ToString())
		intersections += stepIntersections
		points = append(points, point)
	}
	return points, intersections
}
