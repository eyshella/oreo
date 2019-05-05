package calculation

import (
	"log"
	"math/rand"
	"oreo/models"
	"time"
)

func Step(start models.Point, polygon models.Polygon, way int, alpha float64) (models.Point, int) {
	log.Printf("Step started. Params: start: %s, polygon: %s, way: %d, alpha: %f", start.ToString(), polygon.ToString(), way, alpha)
	verticleNumber := way
	currentDirection := models.NewVector(start, polygon.Vertices[way])
	currentDistance := currentDirection.Length()
	remaining := currentDistance * alpha
	intersections := 0
	isReflectionFromVerticle := true
	sectionFrom := -1
	sectionTo := -1
	for remaining > 0 {
		log.Printf("Step. Iteration params: verticleNumber: %d, currentDirection: %s, currentDistance: %f, remaining: %f, intersections: %d, isReflectionFromVerticle: %t, sectionFrom: %d, sectionTo: %d", verticleNumber, currentDirection.ToString(), currentDistance, remaining, intersections, isReflectionFromVerticle, sectionFrom, sectionTo)
		if remaining <= currentDistance {
			log.Printf("Step. Remaining distance <= current distance")
			currentDirection = *currentDirection.MultiplyOnScalar(remaining / currentDistance)
			remaining = 0
		} else {
			log.Printf("Step. Remaining distance > current distance")
			intersections++
			remaining -= currentDistance
			var newDirection *models.Vector
			if isReflectionFromVerticle {
				log.Printf("Step. Reflecting from verticle.")
				newDirection = ReflectFromPolygonVertice(currentDirection, polygon, verticleNumber)
			} else {
				log.Printf("Step. Reflecting from section.")
				newDirection = ReflectFromSection(currentDirection, models.NewVector(polygon.Vertices[sectionFrom], polygon.Vertices[sectionTo]))
			}
			from := currentDirection.To
			intersection, _sectionFrom, _sectionTo := FindIntersectionRayAndPolygon(*newDirection, polygon)
			sectionFrom = _sectionFrom
			sectionTo = _sectionTo
			currentDirection = models.NewVector(from, *intersection)
			currentDistance = currentDirection.Length()
			isReflectionFromVerticle = false
			if polygon.Vertices[sectionFrom].Equal(currentDirection.To) {
				isReflectionFromVerticle = true
				verticleNumber = sectionFrom
			}
			if polygon.Vertices[sectionTo].Equal(currentDirection.To) {
				isReflectionFromVerticle = true
				verticleNumber = sectionTo
			}
		}
	}
	log.Printf("Step finished. Return: point: %s, number of intersections: %d", currentDirection.To.ToString(), intersections)
	return currentDirection.To, intersections
}

func Walk(start models.Point, polygon models.Polygon, steps int) ([]models.Point, int) {
	log.Printf("Walk started. Params: start: %s, polygon: %s, steps: %d", start.ToString(), polygon.ToString(), steps)
	rand.Seed(time.Now().UnixNano())
	point := start
	points := make([]models.Point, 0)
	points = append(points, point)
	intersections := 0
	for i := 0; i < steps; i++ {
		alpha := 100.0
		way := rand.Intn(len(polygon.Vertices))
		log.Printf("Walk. Starting %dth step.", i+1)
		point, stepIntersections := Step(point, polygon, way, alpha)
		log.Printf("Walk. Finished %dth step. Point: %s, number of intersections: %d", i+1, point.ToString(), stepIntersections)
		intersections += stepIntersections
		points = append(points, point)
	}
	log.Printf("Walk finished. Number of points:%d, number of intersections: %d", len(points), intersections)
	return points, intersections
}
