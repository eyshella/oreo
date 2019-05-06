package models

import (
	"fmt"
	"oreo/config"
)

type Point struct {
	X float64
	Y float64
}

func NewPoint(x float64, y float64) Point {
	return Point{
		X: x,
		Y: y,
	}
}

func (p *Point) Add(c Point) {
	p.X += c.X
	p.Y += c.Y
}

func (p *Point) Subtract(c Point) {
	p.X -= c.X
	p.Y -= c.Y
}

func (p Point) Equal(c Point) bool {
	return NewVector(p, c).Length() < config.AppConfig.Accuracy
}

func (p Point) ToString() string {
	return fmt.Sprintf("X: %f, Y: %f", p.X, p.Y)
}
