package struct_entity

type ColorPoint struct { //成员*Point是另一个结构体的指针，和Point结构体是组合关系
	*Point
	Color string
}

type Point struct {
	X, Y int
}

func (p Point) Distance() int {
	return p.X - p.Y
}
