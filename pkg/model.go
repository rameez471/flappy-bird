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
	sub      chan struct{}
	response int
}

func InitialModel() model {
	return model{
		sub:      make(chan struct{}),
		response: 0,
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(
		listenForActivity(m.sub),
		waitForActivity(m.sub),
	)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case tea.KeyMsg:
		return m, tea.Quit
	case responseMsg:
		m.response++
		return m, waitForActivity(m.sub)
	}
	return m, nil
}

func (m model) View() string {
	s := fmt.Sprintf("\n Event recieved: %d\n", m.response)
	return s
}
