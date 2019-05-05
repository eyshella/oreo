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
	polygon := models.NewRegPolygon(6)
	calculation.Walk(models.NewPoint(0, 0), polygon, 100)
	fmt.Println("Press Enter to close")
	helpers.ReadFromConsole()
	log.Println("App finished!")
}
