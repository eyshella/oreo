package helpers

import (
	"encoding/csv"
	"fmt"
	"log"
	"oreo/models"
	"os"
)

func WriteResultCsvFile(path string, points []models.Point) {
	log.Printf("WriteResultCsvFile started. Path: %s", path)
	result := make([][]string, 0)
	result = append(result, []string{"X", "Y"})
	for _, point := range points {
		pointStr := make([]string, 2)
		pointStr[0] = fmt.Sprintf("%f", point.X)
		pointStr[1] = fmt.Sprintf("%f", point.Y)
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
