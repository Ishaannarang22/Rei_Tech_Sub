package handlers

import (
	"net/http"
	"reitechsub/pkg/authentication"
	"reitechsub/pkg/utilsdb"
)

func (h *Handler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	// incase the user still hasn't submitted credentials
	if r.Method != http.MethodPost {
		h.Tmpl.Execute(w, nil)
	} else {
		// Get the POST data
		userIn := utilsdb.User{
			Username: r.FormValue("UsernameInput"),
			Password: r.FormValue("PasswordInput"),
		}

		// Read from the DB
		userDB, err := utilsdb.ReadUser(userIn)
		if err != nil {
			h.Tmpl.Execute(w, nil)
		}

		// If the user doesn't exist
		if userDB.Password == "" {
			h.Tmpl.Execute(w, nil)
		}

		// Check Password
		passBool, _ := authentication.CheckPass(userDB.Password, userIn.Password)
		if passBool != true {
			h.Tmpl.Execute(w, nil)
		}

		// Add cookie to client
		cookie := http.Cookie{Name: "tkn-key", Value: userDB.Token, Secure: true}
		http.SetCookie(w, &cookie)

		// Return a redirect
		redirectURL := "http://0.0.0.0:80/u/" + userIn.Username + "/dashboard/" + r.FormValue("StatusInput")
		http.Redirect(w, r, redirectURL, http.StatusSeeOther)
	}
}
