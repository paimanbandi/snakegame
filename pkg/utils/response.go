package utils

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"net/http"
)

func GenerateSecureID() string {
	bytes := make([]byte, 16) // 16 bytes = 128 bits, similar to a UUID
	if _, err := rand.Read(bytes); err != nil {
		panic(err) // Handle error
	}
	return hex.EncodeToString(bytes)
}

func MakeResponse(w http.ResponseWriter, r *http.Request, res any) {
	if w.Header().Get("Content-Type") != "" {
		return // if headers already set due to an error no need to write again
	}
	// Set the content type to application/json
	w.Header().Set("Content-Type", "application/json")

	// Marshal the struct into JSON
	jsonResponse, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error marshaling JSON"))
		return
	}

	// Write the JSON response
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
