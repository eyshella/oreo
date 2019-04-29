package main

import (
	"log"
	"oreo/models"
)

func main() {
	log.Println("App started!")
	polygon := models.NewPolygon(10)
	for _, v := range polygon.Vertices {
		log.Printf("X:%f Y:%f", v.X, v.Y)
	}
	log.Println("App finished!")
}
