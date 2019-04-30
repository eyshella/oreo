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
	/* polygon := models.NewRegPolygon(10)
	for _, v := range polygon.Vertices {
		log.Printf("X:%f Y:%f", v.X, v.Y)
	} */
	point := calculation.FindIntersectionOfRayAndSection(models.Vector{From: models.Point{X: 0, Y: 0}, To: models.Point{X: 1, Y: 0}}, models.Vector{From: models.Point{X: 1, Y: 0}, To: models.Point{X: 0, Y: 1}})
	log.Printf("X:%f Y:%f", point.X, point.Y)
	fmt.Println("Press Enter to close")
	helpers.ReadFromConsole()
	log.Println("App finished!")
}
