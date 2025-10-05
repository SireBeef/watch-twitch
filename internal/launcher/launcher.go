package launcher

import (
	"fmt"
	"os"
	"os/exec"
)

type StreamLauncher struct {
	browserToken string
}

func NewStreamLauncher(browserToken string) *StreamLauncher {
	return &StreamLauncher{
		browserToken: browserToken,
	}
}

func (sl *StreamLauncher) Launch(name string) {
	command := fmt.Sprintf(`streamlink --twitch-low-latency --twitch-disable-ads -p mpv --player-args "--gpu-context=wayland --ontop" https://twitch.tv/%s best > /dev/null 2>&1 &`, name)

	if sl.browserToken != "" {
		command = fmt.Sprintf(`streamlink --twitch-low-latency "--twitch-api-header=Authorization=OAuth %s" -p mpv --player-args "--gpu-context=wayland --ontop"  https://twitch.tv/%s best > /dev/null 2>&1 &`, sl.browserToken, name)
	}

	cmd := exec.Command("bash", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		fmt.Println("error:", err)
	}
}

type ChatLauncher struct{}

func NewChatLauncher() *ChatLauncher {
	return &ChatLauncher{}
}

func (cl *ChatLauncher) Launch(name string) {
	cmd := exec.Command("bash", "-c",
		fmt.Sprintf(`chatterino -c %s 2>&1 &`, name),
	)
	err := cmd.Start()
	if err != nil {
		fmt.Println("Failed to open Chatterino:", err)
	}
}
