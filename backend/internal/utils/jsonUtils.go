package utils

import (
	"encoding/json"
	"fmt"
	"github.com/KainoaGardner/csc/internal/types"
	"net/http"
)

func ParseJSON(r *http.Request, payload any) error {
	if r.Body == nil {
		return fmt.Errorf("Missing request body")

	}
	return json.NewDecoder(r.Body).Decode(payload)
}

func ParseMsgJSON[T any](msg types.IncomingMessage) (T, error) {
	var result T

	if msg.Data == nil {
		return result, fmt.Errorf("Missing request data")

	}

	jsonData, err := json.Marshal(msg.Data)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJSON(w, status, map[string]string{"error": err.Error()})
}

func WriteResponse(w http.ResponseWriter, status int, message string, data any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	var response types.APIRespone
	response.Message = message
	response.Data = data
	return json.NewEncoder(w).Encode(response)
}
