package models

// Streamer represents a Twitch streamer with their current stream information
type Streamer struct {
	Name    string
	Content string
}

func (s Streamer) Title() string       { return s.Name }
func (s Streamer) Description() string { return s.Content }
func (s Streamer) FilterValue() string { return s.Name }
