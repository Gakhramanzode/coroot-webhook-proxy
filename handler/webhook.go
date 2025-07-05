package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Gakhramanzode/coroot-webhook-proxy/config"
	"github.com/Gakhramanzode/coroot-webhook-proxy/vk"
)

type CorootWebhook struct {
	Status      string   `json:"status"`
	Application string   `json:"application"`
	Summary     []string `json:"summary"`
	URL         string   `json:"url"`
}

func WebhookHandler(cfg config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
			return
		}

		var payload CorootWebhook
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			log.Println("Invalid JSON:", err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		// Logging received data JSON
		log.Printf("Received webhook from Coroot: %+v", payload)

		// Validating required fields
		if cfg.IgnoreEmptySummary &&
			len(payload.Summary) == 1 &&
			payload.Summary[0] == "No notable changes" {
			log.Println("Ignored: summary contains only 'No notable changes'")
			w.WriteHeader(http.StatusNoContent)
			return
		}

		// Building message
		text := fmt.Sprintf("%s: %s\n", payload.Status, payload.Application)
		for _, line := range payload.Summary {
			text += "- " + line + "\n"
		}
		text += "\n" + payload.URL

		if err := vk.SendMessage(cfg, text); err != nil {
			log.Println("Failed to send message to VK:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))

	}
}
