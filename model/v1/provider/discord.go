package model

import "time"

type DiscordMessage struct {
	ID        string    `json:"id"`
	Author    string    `json:"author"`
	Content   string    `json:"content"`
	Timestamp time.Time `json:"timestamp"`
	Period    string    `json:"period"`
}
