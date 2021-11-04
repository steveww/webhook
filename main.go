package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

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

	// dump the request
	log.Println(">>>New Request<<<")
	for key, val := range r.Header {
		log.Println(key, " : ", val)
	}

	var prettyJSON bytes.Buffer
	buffer := new(bytes.Buffer)
	buffer.ReadFrom(r.Body)
	err := json.Indent(&prettyJSON, buffer.Bytes(), "", "\t")
	if err != nil {
		log.Println("JSON parse error", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	log.Println(string(prettyJSON.Bytes()))
}

func main() {
	log.Println("Server starting on 8080")
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/health", handleHealth)
	http.HandleFunc("/ready", handleReady)
	http.HandleFunc("/webhook", handleWebhook)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
