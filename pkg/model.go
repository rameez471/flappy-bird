package pkg

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
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
	sub      chan struct{}
	response int
	screen   *Screen
	bird     *Bird
}

func InitialModel() model {
	return model{
		sub:      make(chan struct{}),
		response: 0,
		screen:   NewScreen(),
		bird:     NewBird(),
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(
		listenForActivity(m.sub),
		waitForActivity(m.sub),
		tea.EnterAltScreen,
	)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return m, tea.Quit
	case tea.WindowSizeMsg:
		m.screen.Update(msg.Width, msg.Height)
	case responseMsg:
		m.response++
		return m, waitForActivity(m.sub)
	}
	return m, nil
}

func (m model) View() string {
	//width, height := m.screen.GetDim()
	birdX, birdY := m.bird.Position.GetPoint()
	var birdStyle = lipgloss.NewStyle().PaddingTop(birdY).PaddingLeft(birdX)
	s := birdStyle.Render("@")

	return s
}
