package models

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
