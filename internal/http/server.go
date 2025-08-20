package http

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"yourproject/configs"
	"yourproject/internal/user"

	"gorm.io/gorm"
)

func StartServer(cfg *configs.Config, db *gorm.DB) error {
	// Initialize repositories
	userRepo := user.NewUserRepository(db)

	// Initialize handlers
	userHandler := user.NewUserHandler(userRepo)

	// Setup router
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)

	// Health check route
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"status":"ok"}`))
	})

	// User routes
	r.Route("/users", func(r chi.Router) {
		r.Post("/", userHandler.CreateUser)
		r.Get("/{id}", userHandler.GetUser)
	})

	// Start server
	addr := ":" + cfg.ServerPort
	log.Printf("Server listening on %s", addr)
	return http.ListenAndServe(addr, r)
}
