package helpers

import (
	"encoding/csv"
	"fmt"
	"log"
	"oreo/models"
	"os"
)

func WriteResultCsvFile(path string, points []models.Point, alphas []float64, intersections []int) {
	log.Printf("WriteResultCsvFile started. Path: %s", path)
	result := make([][]string, 0)
	result = append(result, []string{"X", "Y", "Alpha", "Intersections"})
	for i, point := range points {
		pointStr := make([]string, 4)
		pointStr[0] = fmt.Sprintf("%f", point.X)
		pointStr[1] = fmt.Sprintf("%f", point.Y)
		pointStr[2] = fmt.Sprintf("%f", alphas[i])
		pointStr[3] = fmt.Sprintf("%d", intersections[i])
		result = append(result, pointStr)
	}
	file, err := os.Create(path)
	if err != nil {
		log.Fatalf("WriteResultCsvFile Error: %s", err.Error())
		return
	}
	writer := csv.NewWriter(file)
	err = writer.WriteAll(result)
	if err != nil {
		log.Fatalf("WriteResultCsvFile Error: %s", err.Error())
		return
	}
	log.Printf("WriteResultCsvFile finished.")
}
