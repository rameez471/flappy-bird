package pkg

type Screen struct {
	width  int
	height int
}

func NewScreen() *Screen {
	return &Screen{
		width:  0,
		height: 0,
	}
}

func (s *Screen) Update(x, y int) {
	s.width = x
	s.height = y
}

func (s *Screen) GetDim() []int {
	dim := []int{s.width, s.height}
	return dim
}
