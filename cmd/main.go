package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"refactoring/internal/config"
	"refactoring/internal/handlers"
	"refactoring/internal/storage"
	"time"
)

const nameConfig = "config.yaml"

func main() {
	cfg, err := config.GetConfigFromFile(nameConfig)
	if err != nil {
		log.Fatal("unable to get config file name from env")
	}
	file, err := storage.New(cfg)
	if err != nil {
		log.Fatal("failed to open storage file")
	}
	userApiHandle := handlers.NewUserApiHandler(*file, cfg)

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	//r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(time.Now().String()))
	})

	r.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Route("/users", func(r chi.Router) {
				r.Get("/search", userApiHandle.SearchUsers)
				r.Post("/create", userApiHandle.CreateUser)

				r.Route("/{id}", func(r chi.Router) {
					r.Get("/get", userApiHandle.GetUser)
					r.Patch("/update", userApiHandle.UpdateUser)
					r.Delete("/delete", userApiHandle.DeleteUser)
				})
			})
		})
	})

	http.ListenAndServe(":3333", r)
}
