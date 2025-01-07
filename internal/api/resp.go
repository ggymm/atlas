package api

import (
	"encoding/json"
	"net/http"
)

type Api struct {
}

type Result struct {
	Msg     string `json:"msg"`
	Data    any    `json:"data"`
	Success bool   `json:"success"`
}

func (h *Api) ok(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(Result{
		Msg:     "",
		Data:    data,
		Success: true,
	})
	if err != nil {
		internalServerError(w)
		return
	}
}

func (h *Api) error(w http.ResponseWriter, code int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(Result{
		Msg:     msg,
		Data:    nil,
		Success: false,
	})
	if err != nil {
		internalServerError(w)
		return
	}
}

func badRequest(w http.ResponseWriter) {
	http.Error(w, "Bad Request", http.StatusBadRequest)
}

func methodNotAllowed(w http.ResponseWriter) {
	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
}

func internalServerError(w http.ResponseWriter) {
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
}
