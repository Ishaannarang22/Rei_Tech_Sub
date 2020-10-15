package handlers

import (
	"net/http"
)

func (h *Handler) HomePage(w http.ResponseWriter, r *http.Request) {
	h.Tmpl.Execute(w, nil)
}
