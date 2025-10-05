package services

import (
	"log"

	"github.com/charmbracelet/bubbles/list"
	"github.com/nicklaw5/helix"
	"watch-twitch/internal/models"
)

// TwitchService handles interactions with the Twitch API
type TwitchService struct {
	clientID        string
	userAccessToken string
	userID          string
}

func NewTwitchService(clientID, userAccessToken, userID string) *TwitchService {
	return &TwitchService{
		clientID:        clientID,
		userAccessToken: userAccessToken,
		userID:          userID,
	}
}

func (ts *TwitchService) GetLiveFollowed() []list.Item {
	client, _ := helix.NewClient(&helix.Options{
		ClientID:        ts.clientID,
		UserAccessToken: ts.userAccessToken,
	})

	resp, err := client.GetFollowedStream(&helix.FollowedStreamsParams{
		UserID: ts.userID,
	})

	if err != nil {
		log.Fatal(err)
	}

	items := []list.Item{}
	for _, stream := range resp.Data.Streams {
		items = append(items, models.Streamer{Name: stream.UserName, Content: stream.GameName})
	}
	return items
}
