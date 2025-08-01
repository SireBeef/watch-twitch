package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/joho/godotenv"
	"github.com/nicklaw5/helix"
)

type streamer struct {
	name    string
	content string
}

func (s streamer) Title() string       { return s.name }
func (s streamer) Description() string { return s.content }
func (s streamer) FilterValue() string { return s.name }

type model struct {
	list         list.Model
	promptActive bool
	selected     streamer
}

type launchMode int

const (
	modeNone launchMode = iota
	modeStream
	modeChat
	modeBoth
)

func getLiveFollowed(clientID, userAccessToken string, userID string) []list.Item {
	client, _ := helix.NewClient(&helix.Options{
		ClientID:        clientID,
		UserAccessToken: userAccessToken,
	})

	resp, err := client.GetFollowedStream(&helix.FollowedStreamsParams{
		UserID: userID,
	})

	if err != nil {
		log.Fatal(err)
	}

	items := []list.Item{}
	for _, stream := range resp.Data.Streams {
		items = append(items, streamer{name: stream.UserName, content: stream.GameName})
	}
	return items
}

func launchStream(name string) {
	command := fmt.Sprintf(`streamlink --twitch-low-latency --twitch-disable-ads -p mpv --player-args "--gpu-context=wayland --ontop" https://twitch.tv/%s best > /dev/null 2>&1 &`, name)

	browserToken := os.Getenv("BROWSER_AUTH_TOKEN")

	if browserToken != "" {
		command = fmt.Sprintf(`streamlink --twitch-low-latency "--twitch-api-header=Authorization=OAuth %s" -p mpv --player-args "--gpu-context=wayland --ontop"  https://twitch.tv/%s best > /dev/null 2>&1 &`, browserToken, name)
	}

	cmd := exec.Command("bash", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		fmt.Println("error:", err)
	}
}

func promptForChat(name string) {
	cmd := exec.Command("bash", "-c",
		fmt.Sprintf(`chatterino -c %s 2>&1 &`, name),
	)
	err := cmd.Start()
	if err != nil {
		fmt.Println("Failed to open Chatterino:", err)
	}
}

func initModel(items []list.Item) model {
	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.Title = "Live Twitch Streamers"
	return model{list: l}
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			if m.promptActive {
				// Already prompting, do nothing
				return m, nil
			}
			item := m.list.SelectedItem().(streamer)
			m.selected = item
			m.promptActive = true
			return m, nil

		case "1":
			if m.promptActive {
				launchStream(m.selected.name)
				return m, tea.Quit
			}
		case "2":
			if m.promptActive {
				promptForChat(m.selected.name)
				return m, tea.Quit
			}
		case "3":
			if m.promptActive {
				launchStream(m.selected.name)
				promptForChat(m.selected.name)
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

func (m model) View() string {
	if m.promptActive {
		return fmt.Sprintf(
			"What do you want to launch for %s?\n[1] Stream\n[2] Chat\n[3] Both\n\nPress 1, 2, or 3...",
			m.selected.name,
		)
	}
	return m.list.View()
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	clientID := os.Getenv("CLIENT_ID")
	userID := os.Getenv("USER_ID")
	userAccessToken := os.Getenv("USER_ACCESS_TOKEN")

	items := getLiveFollowed(clientID, userAccessToken, userID)

	if len(items) == 0 {
		fmt.Println("No live followed streamers.")
		os.Exit(0)
	}

	p := tea.NewProgram(initModel(items))
	if _, err := p.Run(); err != nil {
		fmt.Println("error running program:", err)
		os.Exit(1)
	}
}
