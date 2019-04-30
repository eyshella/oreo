package models

import "math"

type Vector struct {
	From Point
	To   Point
}

func NewVector(from Point, to Point) Vector {
	return Vector{
		From: from,
		To:   to,
	}
}

func (v Vector) Length() float64 {
	return math.Sqrt((v.From.X-v.To.X)*(v.From.X-v.To.X) + (v.From.Y-v.To.Y)*(v.From.Y-v.To.Y))
}

func (v Vector) Add(c Vector) *Vector {
	v.To = Point{
		X: c.To.X - c.From.X + v.To.X,
		Y: c.To.Y - c.From.Y + v.To.Y,
	}
	return &v
}

func (v Vector) Subtract(c Vector) *Vector {
	v.From = Point{
		X: c.To.X - c.From.X + v.From.X,
		Y: c.To.Y - c.From.Y + v.From.Y,
	}
	return &v
}

func (v Vector) Multiply(c float64) *Vector {
	if v.Length() == 0 {
		return nil
	}
	v.To.X = c*(v.To.X-v.From.X) + v.From.X
	v.To.Y = c*(v.To.Y-v.From.Y) + v.From.Y
	return &v
}
