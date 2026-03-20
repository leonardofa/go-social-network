package response

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSON sends a JSON response with the given HTTP status code and data payload to the http.ResponseWriter.
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			log.Fatal(err)
		}
	}
}

// ERROR sends an error response with the given HTTP status code and error message to the http.ResponseWriter.
func ERROR(w http.ResponseWriter, statusCode int, err error) {
	JSON(w, statusCode, struct {
		Error string `json:"error"`
	}{Error: err.Error()})
}
