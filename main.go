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
	polygon := models.NewRegPolygon(4)
	for _, v := range polygon.Vertices {
		log.Printf("X:%f Y:%f", v.X, v.Y)
	}
	reflection := calculation.ReflectFromPolygonVertice(models.Vector{From: models.Point{X: 0, Y: 0}, To: models.Point{X: 1, Y: 0}}, polygon, 0)
	log.Printf("From: X:%f Y:%f ; To: X:%f Y:%f", reflection.From.X, reflection.From.Y, reflection.To.X, reflection.To.Y)
	fmt.Println("Press Enter to close")
	helpers.ReadFromConsole()
	log.Println("App finished!")
}
