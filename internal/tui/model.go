package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"watch-twitch/internal/launcher"
	"watch-twitch/internal/models"
)

type Model struct {
	list           list.Model
	showingPrompt  bool
	selected       models.Streamer
	streamLauncher *launcher.StreamLauncher
	chatLauncher   *launcher.ChatLauncher
}

func NewModel(items []list.Item, streamLauncher *launcher.StreamLauncher, chatLauncher *launcher.ChatLauncher) Model {
	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.Title = "Live Twitch Streamers"
	return Model{
		list:           l,
		streamLauncher: streamLauncher,
		chatLauncher:   chatLauncher,
	}
}

func (m Model) Init() tea.Cmd { return nil }

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			if m.showingPrompt {
				return m, nil
			}
			item := m.list.SelectedItem().(models.Streamer)
			m.selected = item
			m.showingPrompt = true
			return m, nil

		case "1":
			if m.showingPrompt {
				m.streamLauncher.Launch(m.selected.Name)
				return m, tea.Quit
			}
		case "2":
			if m.showingPrompt {
				m.chatLauncher.Launch(m.selected.Name)
				return m, tea.Quit
			}
		case "3":
			if m.showingPrompt {
				m.streamLauncher.Launch(m.selected.Name)
				m.chatLauncher.Launch(m.selected.Name)
				return m, tea.Quit
			}
		case "ctrl+c", "q":
			return m, tea.Quit
		}

	case tea.WindowSizeMsg:
		m.list.SetSize(msg.Width, msg.Height)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	if m.showingPrompt {
		return fmt.Sprintf(
			"What do you want to launch for %s?\n[1] Stream\n[2] Chat\n[3] Both\n\nPress 1, 2, or 3...",
			m.selected.Name,
		)
	}
	return m.list.View()
}
