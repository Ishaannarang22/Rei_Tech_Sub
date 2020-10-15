package handlers

import (
	"html/template"
)

// This allows websites to have html pages
type Handler struct {
	Tmpl *template.Template
}

func NewHandler(tmpl *template.Template) *Handler {
	return &Handler {
		Tmpl: tmpl,
	}
}

