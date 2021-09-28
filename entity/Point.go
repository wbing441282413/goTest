package entity

type ColorPoint struct {
	*Point
	Color string
}

type Point struct {
	X, Y int
}

func (p Point) Distance() int {
	return p.X - p.Y
}
