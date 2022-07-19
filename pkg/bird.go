package pkg

import (
	"github.com/charmbracelet/lipgloss"
)

const (
	posX float64 = 6
	posY float64 = 4
)

type Bird struct {
	Position *Point
	Vel      *Vector
	Acc      *Vector
}

func NewBird() *Bird {
	return &Bird{
		Position: NewPoint(6, 4),
		Vel:      NewVector(0, 2),
		Acc:      NewVector(0, 1.2),
	}
}

func (b *Bird) Move(delta float64) {
	b.Vel.Mul(b.Acc, delta)
	b.Position.Add(b.Vel, delta)
}

func (b *Bird) Render() string {
	var style = lipgloss.NewStyle().PaddingTop(int(b.Position.Y)).PaddingLeft(int(b.Position.X))
	s := style.Render("@")
	return s
}

func (b *Bird) Jump(delta float64) {
	b.Vel.Y -= 2
	b.Vel.Mul(b.Acc, delta)
	b.Position.Add(b.Vel, delta)
}
