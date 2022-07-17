package pkg

const (
	posX = 6
	posY = 4
)

type Bird struct {
	Position Point
	Vel      Vector
	Acc      Vector
}

func NewBird() *Bird {
	return &Bird{
		Position: NewPoint(posX, posY),
	}
}

func (b *Bird) setPotion(p Point) {
	b.Position = p
}
