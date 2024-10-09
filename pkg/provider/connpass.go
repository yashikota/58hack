package provider

import (
	"encoding/base64"
	"io"
	"log/slog"
	"net/http"
	"os"
)

func ConnpassProvider() ([]string, error) {
	userID := os.Getenv("CONNPASS_USER_ID")
	if userID == "" {
		slog.Warn("Connpass : CONNPASS_USER_ID environment variable is not set")
		return []string{}, nil
	}

	pass := os.Getenv("CONNPASS_PASS")
	if pass == "" {
		slog.Warn("Connpass : CONNPASS_PASS environment variable is not set")
		return []string{}, nil
	}

	auth := userID + ":" + pass
	authHeader := "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))

	url := "https://chronotes.yashikota.com/connpass/api/v1/event/?keyword=python"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		slog.Error("Connpass : Failed to create request")
		return []string{}, err
	}

	req.Header.Set("Authorization", authHeader)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		slog.Error("Connpass : Failed to send request")
		return []string{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Error("Connpass : Failed to read response body")
		return []string{}, err
	}
	return []string{string(body)}, nil
}
