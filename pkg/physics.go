package pkg

type Point struct {
	X int
	Y int
}

func (p Point) NewPoint(x int, y int) *Point {
	return &Point{
		X: x,
		Y: y,
	}
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
