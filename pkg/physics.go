package pkg

type Point struct {
	X float64
	Y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{
		X: x,
		Y: y,
	}
}

type Vector struct {
	X float64
	Y float64
}

func NewVector(x float64, y float64) *Vector {
	return &Vector{
		X: x,
		Y: y,
	}
}

func (v *Vector) Mul(acc *Vector, delta float64) {
	v.X += acc.X * delta
	v.Y += acc.Y * delta
}

func (p *Point) Add(vel *Vector, delta float64) {
	p.X += vel.X * delta
	p.Y += vel.Y * delta
}
