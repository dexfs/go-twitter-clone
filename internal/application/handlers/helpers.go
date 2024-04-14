package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			log.Fatal(err)
		}
	}
}

func DecodeJSON(r *http.Request, v interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return fmt.Errorf("could not decode JSON: %v", err)
	}
	return nil
}

func JSONError(w http.ResponseWriter, statusCode int, err error) {
	JSON(w, statusCode, struct {
		Erro string `json:"error"`
	}{
		Erro: err.Error(),
	})
}
