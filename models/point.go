package models

type Point struct {
	X float64
	Y float64
}

func (p *Point) Add(c Point) {
	p.X += c.X
	p.Y += c.Y
}

func (p *Point) Subtract(c Point) {
	p.X -= c.X
	p.Y -= c.Y
}
