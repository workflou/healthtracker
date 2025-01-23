package main

import (
	"healthtracker/static"
	"log"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	db := mustNewDatabase(":memory:")

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Group(func(r chi.Router) {
		r.Use(newAuthMiddleware())

		r.Get("/", newHomeHandler(db).Index())

	})
	r.Group(func(r chi.Router) {
		r.Get("/login", newSessionHandler(db).LoginPage())
	})

	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServerFS(static.FS)))

	s := http.Server{
		Addr:    ":4000",
		Handler: r,
	}

	slog.Info("http://localhost:4000")

	log.Fatal(s.ListenAndServe())
}
