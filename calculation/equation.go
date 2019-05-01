package calculation

import "math"

func SolveQuadraticEquation(a, b, c float64) []float64 {
	root := make([]float64, 0)
	if a == 0 && b != 0 {
		root = append(root, c/b)
	}
	if a != 0 {
		d := b*b - 4*a*c
		if d == 0 {
			root = append(root, -b/(2*a))
		} else if d > 0 {
			root = append(root, (-b+math.Sqrt(d))/(2*a))
			root = append(root, (-b-math.Sqrt(d))/(2*a))
		}
	}
	return root
}
