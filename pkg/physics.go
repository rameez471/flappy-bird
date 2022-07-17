package pkg

type Point struct {
	X int
	Y int
}

func NewPoint(x, y int) Point {
	return Point{
		X: x,
		Y: y,
	}
}

func (p *Point) GetPoint() (x, y int) {
	return p.X, p.Y
}

type Vector struct {
	X int
	Y int
}

func (v Vector) NewVector(x int, y int) *Vector {
	return &Vector{
		X: x,
		Y: y,
	}
}
