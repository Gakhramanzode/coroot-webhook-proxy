package config

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	VKURL              string
	ChatID             string
	Token              string
	IgnoreEmptySummary bool
}

func LoadConfig() Config {
	vkURL := os.Getenv("VK_URL")
	chatID := os.Getenv("VK_CHAT_ID")
	token := os.Getenv("VK_TOKEN")
	ignoreEmpty := false

	if val := os.Getenv("IGNORE_EMPTY_SUMMARY"); val != "" {
		parsed, err := strconv.ParseBool(val)
		if err != nil {
			log.Printf("Invalid value for IGNORE_EMPTY_SUMMARY: %s (expected true/false), using false", val)
		} else {
			ignoreEmpty = parsed
		}
	}

	if vkURL == "" || chatID == "" || token == "" {
		log.Fatal("Missing required environment variables: VK_URL, VK_CHAT_ID, VK_TOKEN")
	}

	return Config{
		VKURL:              vkURL,
		ChatID:             chatID,
		Token:              token,
		IgnoreEmptySummary: ignoreEmpty,
	}
}
