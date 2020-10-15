package handlers

import (
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"os"
)

func (h *Handler) HandleMasterDash(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["userhandle"]
	if *h.Username != username {
		w.Write([]byte("Access Denied"))
		return
	}

	h.Tmpl.Execute(w, nil)
}

func HandleVoice(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("data")
	if err != nil {
		panic(err)
	}

	defer file.Close()
	f, err := os.OpenFile("./downloaded", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	io.Copy(f, file)
}
