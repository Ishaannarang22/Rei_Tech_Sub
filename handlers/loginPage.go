package handlers

import (
	"net/http"
	"reitechsub/pkg/authentication"
	"reitechsub/pkg/utilsdb"
	"time"
)

func (h *Handler) StaticLogin(w http.ResponseWriter, r *http.Request) {
	h.Tmpl.Execute(w, nil)
}

func (h *Handler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	// Get the POST data
	userIn := utilsdb.User{
		Username: r.FormValue("UsernameInput"),
		Password: r.FormValue("PasswordInput"),
	}

	// Read from the DB
	userDB, err := utilsdb.ReadUser(userIn)
	if err != nil {
		h.Tmpl.Execute(w, nil)
		return
	}

	// If the user doesn't exist
	if userDB.Password == "" {
		h.Tmpl.Execute(w, nil)
		return
	}

	// Check Password
	passBool, _ := authentication.CheckPass(userDB.Password, userIn.Password)
	if passBool != true {
		h.Tmpl.Execute(w, nil)
		return
	}

	// Add cookie to client
	expiration := time.Now().Add(36500 * 24 * time.Hour)
	cookie := http.Cookie{Name: "tkn-key", Value: userDB.Token, Expires: expiration}
	http.SetCookie(w, &cookie)

	// Return a redirect
	redirectURL := "http://0.0.0.0:80/u/" + userIn.Username + "/dashboard/" + r.FormValue("StatusInput")
	http.Redirect(w, r, redirectURL, http.StatusSeeOther)
}
