package main

import (
	"fmt"
	"log"
	"oreo/calculation"
	"oreo/helpers"
	"oreo/models"
)

func main() {
	log.Println("App started!")
	polygon := models.NewRegPolygon(3)
	for _, v := range polygon.Vertices {
		log.Printf("X:%f Y:%f", v.X, v.Y)
	}
	points, intersections := calculation.Walk(models.NewPoint(0, 0), polygon, 1000)
	log.Printf("Intersection/Steps = %d", intersections/len(points))
	fmt.Println("Press Enter to close")
	helpers.ReadFromConsole()
	log.Println("App finished!")
}
