package api

import (
	"atlas/pkg/data/service"
	"net/http"
)

type VideoApi struct {
	Api
}

func (h *VideoApi) SelectVideos(w http.ResponseWriter, r *http.Request) {
	resp, err := service.SelectVideos(nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.ok(w, resp)
}
