package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
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
			log.Println(pretty_json.String())
		} else {
			log.Println(buffer.String())
		}
	} else {
		log.Println(buffer.String())
	}

	//log.Println(buffer.String())
}

func prettyJSON(ugly bytes.Buffer) (bytes.Buffer, error) {
	var buff bytes.Buffer

	err := json.Indent(&buff, ugly.Bytes(), "", "  ")
	
	return buff, err
}

func main() {
	_, pretty = os.LookupEnv("JSON_PRETTY")

	if pretty {
		log.Println("pretty output enabled")
	}

	log.Println("Server starting on 8080")
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/health", handleHealth)
	http.HandleFunc("/ready", handleReady)
	http.HandleFunc("/webhook", handleWebhook)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
