package response

import (
	"encoding/json"
	"log"
	"net/http"
)

// ApiError api error data representation
type ApiError struct {
	Error string `json:"error"`
}

// JSON handles the json return
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatal(err)
	}
}

// ERROR handles error response
func ERROR(w http.ResponseWriter, r *http.Response) {
	var err ApiError
	json.NewDecoder(r.Body).Decode(&err)
	JSON(w, http.StatusBadRequest, err)
}
