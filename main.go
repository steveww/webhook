package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"net/http"
)

var pretty bool

func handleHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK")
}

func handleReady(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK")
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Webhook Handler")
}

func handleWebhook(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, http.StatusText(http.StatusUnsupportedMediaType), http.StatusUnsupportedMediaType)
		return
	}

	var buffer bytes.Buffer

	// dump the headers
	headers, headersErr := json.Marshal(r.Header)
	
	// read the body
	body := new(bytes.Buffer)
	body.ReadFrom(r.Body)

	if headersErr == nil {
		fmt.Fprintf(&buffer, "{\"headers\": %s, \"body\": %s}", headers, body)
	} else {
		fmt.Fprintf(&buffer, "{\"body\": %s}", body)
	}

	if pretty {
		pretty_json, err := prettyJSON(buffer)
		if err == nil {
			log.Info().RawJSON("payload", pretty_json.Bytes()).Msg("webhook")
		} else {
			log.Info().RawJSON("payload", buffer.Bytes()).Msg("webhook")
		}
	} else {
		log.Info().RawJSON("payload", buffer.Bytes()).Msg("webhook")
	}
}

func prettyJSON(ugly bytes.Buffer) (bytes.Buffer, error) {
	var buff bytes.Buffer

	err := json.Indent(&buff, ugly.Bytes(), "", "  ")
	
	return buff, err
}

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	_, pretty = os.LookupEnv("JSON_PRETTY")

	if pretty {
		log.Info().Msg("pretty output enabled")
	}

	log.Info().Msg("Server starting on 8080")
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/health", handleHealth)
	http.HandleFunc("/ready", handleReady)
	http.HandleFunc("/webhook", handleWebhook)
	log.Fatal().Err(http.ListenAndServe(":8080", nil))
}
