package dto

import (
	"encoding/json"
	"net/http"
)

type Response[T any] struct {
	Message string `json:"message"`
	Data    T      `json:"data"`
	Error   any    `json:"error,omitempty"`
}

func JsonErr(w http.ResponseWriter, status int, msg string, err ...error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	resp := Response[any]{
		Message: msg,
	}
	if len(err) > 0 {
		resp.Error = err[0]
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func JsonOK[T any](w http.ResponseWriter, status int, data T) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	resp := Response[T]{
		Message: "OK",
		Data:    data,
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
