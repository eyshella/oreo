package models

import (
	"fmt"
	"math"
)

type Polygon struct {
	Vertices []Point
}

func NewRegPolygon(n int) Polygon {
	vertices := make([]Point, 0)
	for k := 0; k < n; k++ {
		vertice := Point{
			X: math.Cos(math.Pi * 2 * float64(k) / float64(n)),
			Y: math.Sin(math.Pi * 2 * float64(k) / float64(n)),
		}
		vertices = append(vertices, vertice)
	}
	return Polygon{
		Vertices: vertices,
	}
}

func (p Polygon) ToString() string {
	result := "Vertices:"
	for i, v := range p.Vertices {
		result += fmt.Sprintf(" %dth verticle: %s", i+1, v.ToString())
	}
	return result
}
