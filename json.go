package main

import (
	"encoding/json"
	"log"
	"net/http"
)

/* If an error is provided, it logs the error. For 5XX errors, it logs the message as well. */

func respondWithError(w http.ResponseWriter, code int, msg string, err error) {
	if err != nil {
		log.Println(err) // Log the error if provided
	}
	if code > 499 {
		log.Printf("Responding with 5XX error: %s", msg)
	}
	type errorResponse struct {
		Error string `json:"error"`
	}
	respondWithJSON(w, code, errorResponse{
		Error: msg, // Send the error message in the response
	})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err) // Log the error if JSON marshalling fails
		w.WriteHeader(500) // Send a 500 Internal Server Error status code
		return
	}
	w.WriteHeader(code) // Send the specified HTTP status code
	w.Write(dat) // Write the JSON response
}
