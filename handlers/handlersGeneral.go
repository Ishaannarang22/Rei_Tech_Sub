package handlers

import (
	"html/template"
	"net/http"
	"reitechsub/pkg/utilsdb"
)

// This allows websites to have html pages
type Handler struct {
	Tmpl     *template.Template
	Username *string
}

func NewHandler(tmpl *template.Template) *Handler {
	return &Handler{
		Tmpl: tmpl,
	}
}

// Authentication Middleware
func (h *Handler) AuthUser(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get 'em sugary cookies
		cookie, err := r.Cookie("tkn-key")
		if err != nil {
			http.Redirect(w, r, "http://0.0.0.0:80/login", http.StatusSeeOther)
			return
		}

		// Check tkn from the db
		username, access, err := utilsdb.CheckSession(cookie.Value)
		if err != nil {
			http.Redirect(w, r, "http://0.0.0.0:80/login", http.StatusSeeOther)
			return
		}
		if access != true {
			http.Redirect(w, r, "http://0.0.0.0:80/login", http.StatusSeeOther)
			return
		}

		h.Username = &username
		f(w, r)
	}
}
