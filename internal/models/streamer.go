package models

type Streamer struct {
	Name string
	Game string
}

func (s Streamer) Title() string       { return s.Name }
func (s Streamer) Description() string { return s.Game }
func (s Streamer) FilterValue() string { return s.Name }
