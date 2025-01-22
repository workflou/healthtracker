package main

import (
	"healthtracker/static"
	"log"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	h := newHandler()

	r.Get("/", h.HomePage())
	r.Get("/login", h.LoginPage())

	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServerFS(static.FS)))

	s := http.Server{
		Addr:    ":4000",
		Handler: r,
	}

	slog.Info("http://localhost:4000")

	log.Fatal(s.ListenAndServe())
}
