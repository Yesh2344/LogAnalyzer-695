package utils

import (
    "encoding/json"
    "net/http"
)

// WriteJSON writes a JSON response to the HTTP writer
func WriteJSON(w http.ResponseWriter, data interface{}) {
    json.NewEncoder(w).Encode(data)
}