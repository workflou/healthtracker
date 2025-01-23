package main

import (
	"database/sql"
	"healthtracker/html"
	"html/template"
	"net/http"
)

type sessionHandler struct {
	db *sql.DB
}

func newSessionHandler(db *sql.DB) *sessionHandler {
	return &sessionHandler{
		db: db,
	}
}

func (h *sessionHandler) LoginPage() http.HandlerFunc {
	t := template.Must(template.ParseFS(html.FS, "layout.html", "login.html"))

	return func(w http.ResponseWriter, r *http.Request) {
		t.Execute(w, nil)
	}
}
