package vk

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/Gakhramanzode/coroot-webhook-proxy/config"
)

func SendMessage(cfg config.Config, text string) error {
	v := url.Values{}
	v.Set("text", text)
	v.Set("chatId", cfg.ChatID)
	v.Set("token", cfg.Token)
	v.Set("parseMode", "HTML")

	fullURL := cfg.VKURL + "?" + v.Encode()

	resp, err := http.Get(fullURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	log.Println("VK Teams response:", resp.Status)
	return nil
}
