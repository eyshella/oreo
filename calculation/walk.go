package calculation

import (
	"log"
	"math/rand"
	"oreo/config"
	"oreo/models"
	"strconv"
	"time"
)

func Step(start models.Point, polygon models.Polygon, way int, alpha float64) (result *models.Point, intersections int) {
	log.Printf("Step started. Params: start: %s, polygon: %s, way: %d, alpha: %f", start.ToString(), polygon.ToString(), way, alpha)
	defer func() {
		if r := recover(); r != nil {
			result = nil
			intersections = 0
		}
	}()
	verticleNumber := way
	currentDirection := models.NewVector(start, polygon.Vertices[way])
	currentDistance := currentDirection.Length()
	remaining := currentDistance * alpha
	intersections = 0
	isReflectionFromVerticle := true
	sectionFrom := -1
	sectionTo := -1
	for remaining > 0 {
		log.Printf("Step. Iteration params: verticleNumber: %d, currentDirection: %s, currentDistance: %f, remaining: %f, intersections: %d, isReflectionFromVerticle: %t, sectionFrom: %d, sectionTo: %d", verticleNumber, currentDirection.ToString(), currentDistance, remaining, intersections, isReflectionFromVerticle, sectionFrom, sectionTo)
		if remaining <= currentDistance {
			log.Printf("Step. Remaining distance <= current distance")
			currentDirection = currentDirection.MultiplyOnScalar(remaining / currentDistance)
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
			if intersection == nil {
				log.Printf("Step Error: intersection of ray and polygon is nil for some reason")
				result = nil
				intersections = 0
				log.Printf("Step finished. Return: point: %s, number of intersections: %d", result, intersections)
				return
			}
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
	result = &currentDirection.To
	log.Printf("Step finished. Return: point: %s, number of intersections: %d", result, intersections)
	return
}

func Walk(start models.Point, polygon models.Polygon, steps int) ([]models.Point, []float64, []int) {
	log.Printf("Walk started. Params: start: %s, polygon: %s, steps: %d", start.ToString(), polygon.ToString(), steps)
	rand.Seed(time.Now().UnixNano())
	point := start
	points := make([]models.Point, 0)
	alphas := make([]float64, 0)
	intersections := make([]int, 0)
	for i := 0; i < steps; i++ {
		alpha := GetAlpha()
		alphas = append(alphas, alpha)
		way := rand.Intn(len(polygon.Vertices))
		log.Printf("Walk. Starting %dth step.", i+1)
		stepPoint, stepIntersections := Step(point, polygon, way, alpha)
		if stepPoint == nil {
			log.Printf("Walk. Skipping %dth step. Number of steps will be increased.", i+1)
			steps++
		} else {
			log.Printf("Walk. Finished %dth step. Point: %s, number of intersections: %d", i+1, stepPoint.ToString(), stepIntersections)
			point = *stepPoint
			intersections = append(intersections, stepIntersections)
			points = append(points, *stepPoint)
		}
	}
	globalNumberOfIntersections := 0
	for _, v := range intersections {
		globalNumberOfIntersections += v
	}
	log.Printf("Walk finished. Number of points:%d, number of intersections: %d", len(points), globalNumberOfIntersections)
	return points, alphas, intersections
}

func GetAlpha() float64 {
	log.Printf("GetAlpha started")
	var alpha float64
	if config.Config.Alpha != "Gamma" {
		var err error
		alpha, err = strconv.ParseFloat(config.Config.Alpha, 64)
		if err != nil {
			log.Fatalf("GetAlpha Error: %s", err.Error())
		}
	} else {
		alpha = Gamma(config.Config.GammaShape, config.Config.GammaScale)
	}
	log.Printf("GetAlpha finished. Return: %f", alpha)
	return alpha
}
