package pkg

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type responseMsg struct{}

func listenForActivity(sub chan struct{}) tea.Cmd {
	return func() tea.Msg {
		for {
			time.Sleep(time.Second / Fps)
			sub <- struct{}{}
		}
	}
}

func waitForActivity(sub chan struct{}) tea.Cmd {
	return func() tea.Msg {
		return responseMsg(<-sub)
	}
}

type model struct {
	sub        chan struct{}
	response   int
	screen     *Screen
	bird       *Bird
	lastUpdate time.Time
}

func InitialModel() model {
	return model{
		sub:        make(chan struct{}),
		response:   0,
		screen:     NewScreen(),
		bird:       NewBird(),
		lastUpdate: time.Now(),
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(
		listenForActivity(m.sub),
		waitForActivity(m.sub),
	)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.screen.Update(msg.Width, msg.Height)
	case responseMsg:
		m.response++
		now := time.Now()
		delta := now.Sub(m.lastUpdate).Seconds()
		m.lastUpdate = now
		m.bird.Move(delta)
		return m, waitForActivity(m.sub)
	case tea.KeyMsg:
		switch msg.String() {
		case "enter", " ":
			now := time.Now()
			delta := now.Sub(m.lastUpdate).Seconds()
			m.lastUpdate = now
			m.bird.Jump(delta)
			return m, nil
		case "q":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	dim := m.screen.GetDim()
	s := fmt.Sprintf("(%d, %d) Last Update = %d Velocity= (%f, %f) \n", dim[0], dim[1], m.lastUpdate.Second(), m.bird.Vel.X, m.bird.Vel.Y)
	s += m.bird.Render()
	return s
}
