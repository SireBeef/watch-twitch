package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"watch-twitch/internal/config"
	"watch-twitch/internal/launcher"
	"watch-twitch/internal/services"
	"watch-twitch/internal/tui"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize services with dependency injection
	twitchService := services.NewTwitchService(cfg.ClientID, cfg.UserAccessToken, cfg.UserID)
	streamLauncher := launcher.NewStreamLauncher(cfg.BrowserToken)
	chatLauncher := launcher.NewChatLauncher()

	// Get live followed streamers
	items := twitchService.GetLiveFollowed()

	if len(items) == 0 {
		fmt.Println("No live followed streamers.")
		os.Exit(0)
	}

	// Initialize TUI with injected dependencies
	model := tui.NewModel(items, streamLauncher, chatLauncher)

	// Run the program
	p := tea.NewProgram(model)
	if _, err := p.Run(); err != nil {
		fmt.Println("error running program:", err)
		os.Exit(1)
	}
}
