package vk

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/Gakhramanzode/coroot-webhook-proxy/config"
)

func SendMessage(cfg config.Config, text string) error {
	fullURL := fmt.Sprintf(
		"%s?text=%s&chatId=%s&token=%s",
		cfg.VKURL,
		url.QueryEscape(text),
		url.QueryEscape(cfg.ChatID),
		url.QueryEscape(cfg.Token),
	)

	resp, err := http.Get(fullURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	log.Println("VK Teams response:", resp.Status)
	return nil
}
