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
	/* polygon := models.NewRegPolygon()
	for _, v := range polygon.Vertices {
		log.Printf("X:%f Y:%f", v.X, v.Y)
	}
	point := calculation.FindIntersectionRayAndPolygon(models.Vector{From: models.Point{X: 0, Y: 0}, To: models.Point{X: 1, Y: 1}}, polygon)
	log.Printf("X:%f Y:%f", point.X, point.Y) */
	reflection := calculation.ReflectFromSection(models.Vector{From: models.Point{X: 0, Y: 0}, To: models.Point{X: 1, Y: 1}}, models.Vector{From: models.Point{X: 0, Y: 10}, To: models.Point{X: 1, Y: 0}})
	log.Printf("From: X:%f Y:%f ; To: X:%f Y:%f", reflection.From.X, reflection.From.Y, reflection.To.X, reflection.To.Y)
	fmt.Println("Press Enter to close")
	helpers.ReadFromConsole()
	log.Println("App finished!")
}
