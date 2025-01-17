package api

import (
	"net/http"
)

type TaskApi struct {
	Api
}

func (h *TaskApi) Exec(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		methodNotAllowed(w)
		return
	}
}
