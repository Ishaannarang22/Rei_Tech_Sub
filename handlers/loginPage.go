package handlers

import (
	"net/http"
)

func (h *Handler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	// incase the user still hasn't submitted credentials
	if r.Method!=http.MethodPost {
		h.Tmpl.Execute(w, nil)
	} else {
		// Get the POST data
		// Read from the DB
		// If the user doesn't exist
		// Check Password
		// Create token
		// Add token to DB
		// Add cookie to client
		// Return a redirect
	}
}
