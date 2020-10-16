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
	val := r.FormValue("fname")
	if err != nil {
		panic(err)
	}

	defer file.Close()
	f, err := os.OpenFile("./static/usercache/"+val, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	io.Copy(f, file)
}

func (h *Handler) HandleCacheFile(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["userhandle"]
	if *h.Username != username {
		w.Write([]byte("Access Denied"))
		return
	}

	filename := r.FormValue("identifier")
	newfilename := r.FormValue("date") + "," + r.FormValue("time") + ".wav"

	// Ignore error if directory already exists
	_ = os.Mkdir("./static/user/"+username, 0666)

	// Rename file to appropriate position
	oldLocation := "./static/usercache/" + filename
	newLocation := "./static/user/" + username + "/" + newfilename

	err := os.Rename(oldLocation, newLocation)
	if err != nil {
		panic(err)
	}

	http.Redirect(w, r, "https://0.0.0.0/u/"+username+"/dashboard/master", http.StatusSeeOther)
}
