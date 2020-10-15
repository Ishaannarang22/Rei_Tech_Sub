package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
)

func (h *Handler) HandleMasterDash(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["userhandle"]
	if *h.Username != username {
		w.Write([]byte("Access Denied"))
		return
	}

	h.Tmpl.Execute(w, nil)
}
