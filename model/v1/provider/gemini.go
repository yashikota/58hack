package model

type Gemini struct {
	GitHubUserID     string `json:"github_user_id"`
	DiscordChannelID string `json:"discord_channel_id"`
	SlackChannelID   string `json:"slack_channel_id"`
}

type Response struct {
	Result []string `json:"result"`
	Day    string   `json:"day"`
}
