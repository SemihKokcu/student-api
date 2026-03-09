package main

import (
	"net/http"
	"strconv"
	"student-api/internal/container"
	"student-api/internal/middleware"

	"github.com/go-chi/chi/v5"
	chi_middleware "github.com/go-chi/chi/v5/middleware"
)

func main() {
	app, err := container.New()
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()

	// 1. Global Middleware'ler
	r.Use(middleware.StructuredLogger(app.Logger))
	r.Use(chi_middleware.Recoverer)

	// 2. Rotalar (Route Grouping)
	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/students", func(r chi.Router) {
			r.Post("/", app.StudentHandler.CreateStudent)
			r.Get("/", app.StudentHandler.GetStudent)
			r.Get("/{id}", app.StudentHandler.GetStudent)
			r.Delete("/{id}", app.StudentHandler.DeleteStudent)
		})
	})

	addr := ":" + strconv.Itoa(app.Config.Server.Port)
	app.Logger.Info("Server started", "addr", addr)
	http.ListenAndServe(addr, r)
}
