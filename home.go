package main

import (
	"database/sql"
	"net/http"
)

type homeHandler struct {
	db *sql.DB
}

func newHomeHandler(db *sql.DB) *homeHandler {
	return &homeHandler{
		db: db,
	}
}

func (h *homeHandler) Index() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("home"))
	}
}
