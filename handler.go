package main

import (
	"database/sql"
	"healthtracker/html"
	"html/template"
	"net/http"
)

type handler struct {
	db *sql.DB
}

func newHandler(db *sql.DB) *handler {
	return &handler{
		db: db,
	}
}

func (h *handler) HomePage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("home"))
	}
}

func (h *handler) LoginPage() http.HandlerFunc {
	t := template.Must(template.ParseFS(html.FS, "layout.html", "login.html"))

	return func(w http.ResponseWriter, r *http.Request) {
		t.Execute(w, nil)
	}
}
