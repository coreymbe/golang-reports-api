package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	logInfo("Responding with payload", payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}

func respondWithError(w http.ResponseWriter, err error) {
	logError(err)
	http.Error(w, err.Error(), http.StatusBadRequest)
}

func logError(err error) {
	log.Println(fmt.Sprintf("[ERROR] %v", err))
}

func logInfo(msg string, v interface{}) {
	log.Println(fmt.Sprintf("[INFO] %s: %v", msg, v))
}
