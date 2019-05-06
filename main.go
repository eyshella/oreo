package main

import (
	"log"
	"oreo/calculation"
	"oreo/config"
	"oreo/helpers"
	"oreo/models"
)

func main() {
	config.SetUpConfig("./config.json")
	helpers.SetUpLogsFile("./log.txt")
	log.Println("App started!")
	polygon := models.NewRegPolygon(config.Config.NumberOfVerticles)
	points, _ := calculation.Walk(models.NewPoint(0, 0), polygon, config.Config.Steps)
	helpers.WriteResultCsvFile(config.Config.ResultPath, points)
	log.Println("App finished!")
}
