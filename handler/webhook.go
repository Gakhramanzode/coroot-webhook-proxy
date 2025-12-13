package handler

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/Gakhramanzode/coroot-webhook-proxy/config"
	"github.com/Gakhramanzode/coroot-webhook-proxy/vk"
)

func WebhookHandler(cfg config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
			return
		}

		// read a Coroot UI template text
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println("read body error:", err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
		// debug: restore r.Body (if will need in someplace else)
		r.Body = io.NopCloser(bytes.NewReader(body))

		// logging the final text:
		// text which created via go-tenplate https://docs.coroot.com/alerting/webhook/?example=plain_text
		log.Printf("RAW webhook body (text): %s", string(body))

		// ignoring if message include "No notable changes"
		if cfg.IgnoreEmptySummary && strings.Contains(string(body), "No notable changes") {
			log.Println("Ignored: summary contains only 'No notable changes'")
			w.WriteHeader(http.StatusNoContent)
			return
		}
		// send the final text
		if err := vk.SendMessage(cfg, string(body)); err != nil {
			log.Println("Failed to send message to VK:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// http status OK
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))

	}
}
