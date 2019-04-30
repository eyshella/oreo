package calculation

import "oreo/models"

func Walk(start models.Point, polygon models.Polygon, way int, alpha float64) (models.Point, int) {
	direction := models.NewVector(start, polygon.Vertices[way])
	distance := direction.Length()
	/* path */ _ = distance * alpha

	return start, 0
}
